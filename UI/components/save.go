package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
)

func saveToFile(tempFileName, newFileName string) error {
	// Rename the temp file to the new file name
	return os.Rename(tempFileName, newFileName)
}

func showSaveAsPopup(speakerID, modType, extension string, window fyne.Window) {
	input := widget.NewEntry()
	input.SetPlaceHolder(speakerID + extension)

	var popup *widget.PopUp

	saveBtn := widget.NewButton("Save", func() {
		newFileName := speakerID + "_" + modType + extension
		err := saveToFile("working.tmp", newFileName)
		if err != nil {
			dialog.ShowError(err, window)
		} else {
			dialog.ShowInformation("Success", "File saved successfully", window)
			popup.Hide()
		}
	})

	content := container.NewVBox(
		widget.NewLabel("File already exists. Enter a new file name:"),
		input,
		saveBtn,
	)

	popup = widget.NewModalPopUp(content, window.Canvas())
	popup.Show()
}

func SaveFile(speakerID, modType, extension string, window fyne.Window) {
	fileName := speakerID + "_" + modType + extension
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// File does not exist, rename the working.tmp file
		err := saveToFile("working.tmp", fileName)
		if err != nil {
			dialog.ShowError(err, window)
		} else {
			dialog.ShowInformation("Success", "File saved successfully", window)
			//TODO: Delete working.tmp
		}
	} else {
		// File exists, show popup to get a new file name
		showSaveAsPopup(speakerID, modType, extension, window)
	}
}
