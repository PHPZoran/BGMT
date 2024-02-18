package components

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func saveToFile(tempFileName, newFileName string) error {
	// Rename the temp file to the new file name
	return os.Rename(tempFileName, newFileName)
}

func showSaveAsPopup(creatureID, modType, extension string, currentDirPath string, window fyne.Window) {
	tempFilePath := filepath.Join(currentDirPath, "working.tmp")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter new file name")

	var popup *widget.PopUp

	saveBtn := widget.NewButton("Save", func() {
		userFileName := input.Text
		newFileName := ""
		if userFileName != "" {
			if strings.HasSuffix(userFileName, extension) {
				newFileName = userFileName
			} else {
				newFileName = userFileName + "_" + modType + extension
			}
		} else {
			if creatureID == "" {
				creatureID = "tempCreatureID"
			}
			// Fallback to default naming if no input is provided
			newFileName = creatureID + "_" + modType + extension

		}
		newFilePath := filepath.Join(currentDirPath, newFileName)

		err := saveToFile(tempFilePath, newFilePath)
		if err != nil {
			dialog.ShowError(err, window)
		} else {
			dialog.ShowInformation("Success", "File saved successfully.", window)
			popup.Hide()
		}
	})

	saveBtnBox := container.NewHBox(
		layout.NewSpacer(),
		saveBtn,
		layout.NewSpacer(),
	)

	content := container.NewVBox(
		widget.NewLabel("File already exists or Speaker ID needed. Enter a new file name:"),
		input,
		saveBtnBox,
	)

	popup = widget.NewModalPopUp(content, window.Canvas())
	popup.Show()
}

func SaveFile(creatureID, modType, extension string, currentDirPath string, window fyne.Window) {
	tempFilePath := filepath.Join(currentDirPath, "working.tmp")

	// Check if modType is empty
	if modType == "" {
		log.Printf("modType unset. Currently: " + modType + "\nThis indicates the user is not working on any mods via working.tmp")
		dialog.ShowError(errors.New("you are not in a working mod location. cannot save"), window)
		return
	}

	// Check if creatureID is empty
	if creatureID == "" {
		showSaveAsPopup("INSERT CREATUREID", modType, extension, currentDirPath, window)
		return
	}

	newFileName := creatureID + "_" + modType + extension
	newFilePath := filepath.Join(currentDirPath, newFileName)
	if _, err := os.Stat(newFileName); os.IsNotExist(err) {
		// File does not exist, rename the working.tmp file
		err := saveToFile(tempFilePath, newFilePath)
		if err != nil {
			dialog.ShowError(err, window)
		} else {
			dialog.ShowInformation("Success", "File saved successfully", window)
			deleteErr := os.Remove(tempFilePath)
			if deleteErr != nil {
				if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
					err := errors.New("working.tmp is not present")
					log.Print(err)
				} else {
					err := errors.New("working.tmp is present but could not delete")
					log.Print(err)
				}
			}
		}
	} else {
		// File exists, show popup to get a new file name
		showSaveAsPopup(creatureID, modType, extension, currentDirPath, window)
	}
}
