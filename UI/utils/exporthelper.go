package utils

import (
	"archive/zip"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func WeiDuFileConversion(window fyne.Window) {
	if CreateZipFolder(GetParentDirectory() + "/TestFolder.zip") {
		dialog.ShowInformation("Error", "Error Creating Zip Folder", window)
		return
	}

	WeiDuConversionHelper("dialogue")
	WeiDuConversionHelper("script")
	WeiDuConversionHelper("installation")
	WeiDuConversionHelper("translation")

	dialog.ShowInformation("WeiDu File Status", "WeiDu file conversions are complete.", window)
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
	JSONtoWeiDuHelper(mod, files)
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func CreateZipFolder(filename string) bool {
	fmt.Println(filename)
	if _, err := os.Stat(filename); err != nil {
		_, err := os.Create(filename)
		check(err)

		return false
	}

	return true
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

// ZipFiles compresses one or more files into a zip archive.
// Credit: https://golangcode.com/create-zip-files-in-go/
func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err := addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

// addFileToZip adds a file to the zip archive
func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	// Create a header for the file
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Set the name of the file inside the zip file
	header.Name = filename

	// Add metadata to the file header if needed
	header.Method = zip.Deflate

	// Create a writer for the file in the zip archive
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Write the file data to the zip archive
	_, err = io.Copy(writer, fileToZip)
	return err
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=
