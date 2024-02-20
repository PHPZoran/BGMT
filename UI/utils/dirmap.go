package utils

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

// CreateFileTree creates and returns a tree view of the directory structure starting from dirPath.
func CreateFileTree(dirPath string, onSelect func(string), onDoubleClick func(string)) *widget.Tree {
	dirMap := make(map[string][]string) // Map from parent path to slice of child paths

	var lastSelected string
	var clickTimer *time.Timer
	doubleClickDuration := time.Duration(time.Millisecond) * 100

	var addChildPaths func(string, string)

	// Recursive function to populate dirMap with the directory structure
	addChildPaths = func(parentPath, currentPath string) {
		files, err := os.ReadDir(currentPath)
		if err != nil {
			return
		}
		for _, file := range files {
			// Construct the relative path of the file or subdirectory
			relPath := filepath.Join(parentPath, file.Name())
			if parentPath == "" {
				relPath = file.Name() // Handle root case
			}
			if file.IsDir() {
				// Recursively handle subdirectories; initialize map entry to signify branch
				dirMap[relPath] = []string{}
				addChildPaths(relPath, filepath.Join(currentPath, file.Name()))
			}
			// Append this file or subdirectory to its parent's slice in dirMap
			dirMap[parentPath] = append(dirMap[parentPath], relPath)
		}
	}

	// Initialize directory structure in dirMap
	addChildPaths("", dirPath)

	// Create the tree
	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID { // childUIDs
			return dirMap[id]
		},
		func(id widget.TreeNodeID) bool { // isBranch
			_, exists := dirMap[id]
			return exists
		},
		func(branch bool) fyne.CanvasObject { // create
			return widget.NewLabel("")
		},
		func(id widget.TreeNodeID, branch bool, node fyne.CanvasObject) { // update
			node.(*widget.Label).SetText(filepath.Base(id))
		},
	)

	// Set onSelect behavior
	tree.OnSelected = func(id widget.TreeNodeID) {
		if clickTimer != nil {
			clickTimer.Stop()
		}
		clickTimer = time.AfterFunc(doubleClickDuration, func() {
			if onSelect != nil && lastSelected == id {
				onSelect(id) // Single-click action
			}
		})
		if lastSelected == id {
			// Double-click detected
			clickTimer.Stop()
			if onDoubleClick != nil {
				onDoubleClick(id) // Double-click action
			}
		}
		lastSelected = id
	}

	return tree
}

func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

// CopyDir copies all the content from src location to new location
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcFilePath := path.Join(src, fd.Name())
		dstFilePath := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcFilePath, dstFilePath); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcFilePath, dstFilePath); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// PromptForProjectName prompts the user for a project name and location to create a new project directory.
func PromptForProjectName(window fyne.Window, onCreate func(newPath string)) {
	// Step 1: Prompt for project name
	templatesPath := "templates"
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Project Name")
	nameDialog := dialog.NewCustomConfirm("New Project Name", "Create", "Cancel", entry, func(b bool) {
		if !b {
			return // User cancelled the operation
		}
		projectName := entry.Text

		// Step 2: Select location for the new project
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil || err != nil {
				return // User cancelled the operation or an error occurred
			}
			projectPath := filepath.Join(uri.Path(), projectName)

			// Step 3: Create the directory
			if err := os.Mkdir(projectPath, 0755); err != nil {
				dialog.ShowError(err, window)
				return
			}

			// Step 4: Create indicative BGM project .ini file
			if err := os.WriteFile(projectPath+"/BGM.ini", []byte(""), 0755); err != nil {
				dialog.ShowError(err, window)
				return
			}

			// Step 5: Create associated project folder directories: Dialogue, Script, Installation, Translation
			if err := CopyDir(templatesPath, projectPath); err != nil {
				dialog.ShowError(err, window)
				return
			}

			// Invoke the callback with the new project path
			onCreate(projectPath)
		}, window)
	}, window)
	nameDialog.Show()
}
