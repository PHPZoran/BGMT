package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
)

// CreateFileTree creates and returns a tree view of the directory structure starting from dirPath.
func CreateFileTree(dirPath string, onSelect func(string)) *widget.Tree {
	dirMap := make(map[string][]string) // Map from parent path to slice of child paths
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
		if onSelect != nil {
			onSelect(id)
		}
	}

	return tree
}

// PromptForProjectName prompts the user for a project name and location to create a new project directory.
func PromptForProjectName(window fyne.Window, onCreate func(newPath string)) {
	// Step 1: Prompt for project name
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
			if err := os.Mkdir(projectPath+"/Dialogue", 0755); err != nil {
				dialog.ShowError(err, window)
				return
			}
			if err := os.Mkdir(projectPath+"/Script", 0755); err != nil {
				dialog.ShowError(err, window)
				return
			}
			if err := os.Mkdir(projectPath+"/Installation", 0755); err != nil {
				dialog.ShowError(err, window)
				return
			}
			if err := os.Mkdir(projectPath+"/Translation", 0755); err != nil {
				dialog.ShowError(err, window)
				return
			}

			// Invoke the callback with the new project path
			onCreate(projectPath)
		}, window)
	}, window)
	nameDialog.Show()
}

/*
func PromptForText(window fyne.Window) string {

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Export Folder Name")
	nameDialog := dialog.NewCustomConfirm("Export Project Name", "Create", "Cancel", entry, func(b bool) {
		if !b {
			return
		}
		WeiDuFileConversion(window,entry.Text)
	}, window)
	nameDialog.Show()

	return filename
}*/
