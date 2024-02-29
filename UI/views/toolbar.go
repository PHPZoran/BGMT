package views

import (
	"UI/components"
	"UI/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"path/filepath"
)

// CreateToolbar creates and returns a configured toolbar widget
func CreateToolbar(directoryPath string, window fyne.Window, creatureID string, modType string, extension string, currentDirectoryPath string) *fyne.Container {
	//setting filepath variables
	var currentFilePath = currentDirectoryPath
	if currentFilePath == "" {
		currentFilePath = directoryPath
		log.Println("Current File Path is " + currentFilePath)
	}
	workingFilePath := filepath.Join(currentFilePath, "working.tmp")

	//homeButton checks for unsaved working.tmp file before navigating back to home.
	homeButton := widget.NewButton("Home", func() {
		if _, err := os.Stat(workingFilePath); err == nil {
			// File exists, show confirmation dialog
			dialog.ShowConfirm("Confirmation", "You have unsaved work. Leaving now will delete your unsaved work. Do you want to continue?",
				func(confirm bool) {
					if confirm {
						// User confirmed, delete the file and navigate
						err := os.Remove(workingFilePath)
						if err != nil {
							log.Printf("Failed to delete unsaved work file '%s': %v", workingFilePath, err)
							dialog.ShowError(fmt.Errorf("failed to delete unsaved work file: %v", err), window)
						}
						NavigateTo(window, directoryPath, MakeHomeView)
					}
				}, window)
		} else {
			NavigateTo(window, directoryPath, MakeHomeView)
		}
	})

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=
	//exportButton TODO
	exportButton := widget.NewButton("Export", func() {
		utils.WeiDuFileConversion(window)
	})

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=
	//saveButton checks for working.tmp file before displaying a save option.
	saveButton := widget.NewButton("Save", func() {
		if _, err := os.Stat(workingFilePath); os.IsNotExist(err) {
			// File does not exist, show error dialog
			log.Printf("Failed to find an unsaved work file '%s':\n%v", workingFilePath, err)
			dialog.ShowError(fmt.Errorf("failed to find an unsaved work file: \n%v", err), window)
		} else {
			log.Printf("Temp file present. Attempting to save.")
			components.SaveFile(creatureID, modType, extension, currentFilePath, window)
		}
	})

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=
	// Create the toolbar using a horizontal container
	toolbar := container.NewHBox(
		homeButton,
		exportButton,
		saveButton,
		layout.NewSpacer(),
	)

	return toolbar
}
