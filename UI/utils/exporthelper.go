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
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkDialog(e error, window fyne.Window) {
	dialog.ShowInformation("Error", e.Error(), window)
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

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
		GPT_Bullshit(zipPath)

		//checkDialog(zipSource(GetInstallationDirectory(), zipPath), window)
		//checkDialog(zipSource(GetTranslationDirectory(), zipPath), window)
		//checkDialog(zipSource(GetScriptDirectory(), zipPath), window)
		//checkDialog(zipSource(GetDialogueDirectory(), zipPath), window)

		//if CreateZipFolder(GetParentDirectory() + "/" + filename + ".zip") {
		//	dialog.ShowInformation("Error", "Error Creating Zip Folder", window)
		//	return
		//}

		//Roadmap for later use
		//WeiDuConversionHelper("dialogue")
		//WeiDuConversionHelper("script")
		//WeiDuConversionHelper("installation")
		//WeiDuConversionHelper("translation")

		dialog.ShowInformation("WeiDu File Status", "WeiDu file conversions are complete and can be found at:\n\n"+
			GetParentDirectory()+"/"+filename, window)

	}, window)
	nameDialog.Show()
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func WeiDuConversionHelper(mod string) {
	var files []string

	//Chain to consolidate all the folders that must be exported
	if mod == "dialogue" {
		directory, err := os.Open(GetDialogueDirectory())
		check(err)
		file, err := directory.ReadDir(-1)
		check(err)

		for _, file := range file {
			files = append(files, GetDialogueDirectory()+"/"+file.Name())
		}
	} else if mod == "script" {
		directory, err := os.Open(GetScriptDirectory())
		check(err)
		file, err := directory.ReadDir(-1)
		check(err)

		for _, file := range file {
			files = append(files, GetScriptDirectory()+"/"+file.Name())
		}
	} else if mod == "installation" {
		directory, err := os.Open(GetInstallationDirectory())
		check(err)
		file, err := directory.ReadDir(-1)
		check(err)

		for _, file := range file {
			files = append(files, GetInstallationDirectory()+"/"+file.Name())
		}
	} else {
		directory, err := os.Open(GetTranslationDirectory())
		check(err)
		file, err := directory.ReadDir(-1)
		check(err)

		for _, file := range file {
			files = append(files, GetTranslationDirectory()+"/"+file.Name())
		}
	}

	fmt.Println(files)
	//JSONtoWeiDuHelper(mod, files)
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func CreateZipFolder(filename string) bool {
	fmt.Println(filename)
	if _, err := os.Stat(filename); err != nil {
		zipFolder, err := os.Create(filename)
		check(err)
		defer zipFolder.Close()

		//filesDialogue := ReadDirectory(GetDialogueDirectory())
		//filesScript := ReadDirectory(GetScriptDirectory())
		//filesInstallation := ReadDirectory(GetInstallationDirectory())
		//filesTranslation := ReadDirectory(GetTranslationDirectory())
		//
		//zipWriter := zip.NewWriter(zipFolder)

		return false
	}

	return true
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func ReadDirectory(directory string) []string {
	var files []string

	dir, err := os.Open(directory)
	check(err)

	file, err := dir.ReadDir(-1)
	check(err)

	for _, file := range file {
		files = append(files, directory+"/"+file.Name())
	}

	return files
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func GPT_Bullshit(zipFolder string) {
	zipFile, err := os.Create(zipFolder)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	directories := []string{GetDialogueDirectory(), GetScriptDirectory(), GetInstallationDirectory(), GetTranslationDirectory()}

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Loop through each directory
	for _, dir := range directories {
		// Walk through the directory and its subdirectories
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Create a new zip file entry
			zipPath, err := filepath.Rel(filepath.Dir(dir), path)
			if err != nil {
				return err
			}

			// Check if the file/directory is not a directory
			if !info.IsDir() {
				// Create a new file entry in the zip writer
				fileWriter, err := zipWriter.Create(zipPath)
				if err != nil {
					return err
				}

				// Open the file to be zipped
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				// Copy the file contents to the zip file
				_, err = io.Copy(fileWriter, file)
				if err != nil {
					return err
				}
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
