package utils

import (
	"archive/zip"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WeiDuFileConversion(window fyne.Window) {
	var filename string
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Export Folder Name")
	nameDialog := dialog.NewCustomConfirm("Export Project Name", "Create", "Cancel", entry, func(b bool) {
		if !b {
			return
		}

		filename = entry.Text
		zipPath := GetParentDirectory() + "/" + filename + ".zip"
		fmt.Println(zipPath)
		exportZip(zipPath)

		dialog.ShowInformation("WeiDu File Status", "WeiDu file conversions are complete and can be found at:\n\n"+
			GetParentDirectory()+"/"+filename, window)

	}, window)
	nameDialog.Show()
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func exportZip(zipFolder string) {
	projectDir := GetParentDirectory()
	projectName := GetModFolder()

	zipFile, err := os.Create(zipFolder)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer func() {
		if err := zipFile.Close(); err != nil {
			fmt.Println("Error closing zip file:", err)
		}
	}()

	directories := []string{GetDialogueDirectory(), GetScriptDirectory(), GetTranslationDirectory(), GetCreatureDirectory()}

	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if err := zipFile.Close(); err != nil {
			fmt.Println("Error closing zip file:", err)
		}
	}()

	// Copy Weidu_Compiler.exe as setup-ProjectName.exe in the zip
	compilerPath := filepath.Join(projectDir, "Weidu_Compiler.exe")
	setupName := fmt.Sprintf("setup-%s.exe", projectName)
	err = addFileToZip(zipWriter, compilerPath, setupName)
	if err != nil {
		fmt.Println("Error adding Weidu_Compiler to zip:", err)
		return
	}
	// Find and add the .tp2 file
	tp2FilePath, err := findTp2File(projectDir)
	if err != nil {
		fmt.Println("Error finding .tp2 file:", err)
		return
	}
	if tp2FilePath != "" { // Check if .tp2 file is found
		tp2ZipPath := filepath.Join(projectName, filepath.Base(tp2FilePath)) // Place inside projectName directory
		err = addFileToZip(zipWriter, tp2FilePath, tp2ZipPath)
		if err != nil {
			fmt.Println("Error adding .tp2 file to zip:", err)
			return
		}
	}

	// Loop through each directory
	for _, dir := range directories {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(projectDir, path)
			if err != nil {
				return err
			}

			// Skip undesired directories and files
			if info.IsDir() {
				// Skip the entire Installation directory
				if strings.Contains(relativePath, "Installation") {
					return filepath.SkipDir
				}
			} else if strings.HasSuffix(info.Name(), ".ini") || strings.HasSuffix(info.Name(), ".txt") || (strings.HasSuffix(info.Name(), ".exe") && info.Name() != "Weidu_Compiler.exe") {
				// Skip specific file types and all .exe files except Weidu_Compiler.exe
				return nil
			}

			// For valid files, add them to the zip
			if !info.IsDir() {
				return addFileToZip(zipWriter, path, filepath.Join(projectName, relativePath))
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error walking through directory:", err)
			return
		}
	}

	fmt.Println("Directories zipped successfully!")
}

// Helper function to add a file to the zip
func addFileToZip(zipWriter *zip.Writer, filePath, zipPath string) error {
	fileWriter, err := zipWriter.Create(zipPath)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	_, err = io.Copy(fileWriter, file)
	return err
}

// Helper function to find the first .tp2 file in a directory
func findTp2File(dirPath string) (string, error) {
	var tp2FilePath string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tp2") {
			tp2FilePath = path // Set the first .tp2 file path found
			return io.EOF      // Use io.EOF to break out of the walk early
		}
		return nil
	})

	if err != nil && err != io.EOF { // io.EOF is expected if a .tp2 file was found
		return "", err
	}
	return tp2FilePath, nil
}
