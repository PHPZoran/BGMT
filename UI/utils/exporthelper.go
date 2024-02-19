package utils

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

		if CreateZipFolder(GetParentDirectory() + "/" + filename + ".zip") {
			dialog.ShowInformation("Error", "Error Creating Zip Folder", window)
			return
		}

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
		_, err := os.Create(filename)
		check(err)

		return false
	}

	return true
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=
