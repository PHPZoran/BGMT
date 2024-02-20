package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"io/ioutil"
	"log"
)

// NavigateTo takes a fyne window and a destination view function and sets the window content to the new view.
func NavigateTo(window fyne.Window, directoryPath string, viewFunc func(string, fyne.Window) fyne.CanvasObject) {
	window.SetContent(viewFunc(directoryPath, window))
	log.Printf("Directory Path is: " + directoryPath + "\n")
}

func HandleErrorAndNavigate(err error, directoryPath, fullPath, selected string, window fyne.Window) { //MakeInstallationView, MakeScriptsView
	if err != nil {
		_, dirErr := ioutil.ReadDir(fullPath)
		if dirErr != nil {
			dialog.ShowError(dirErr, window)
		} else {
			dialog.ShowConfirm("", "Do you want to start a new "+selected+" mod?", func(confirm bool) {
				if confirm {
					logMessage := "Navigating to " + selected + " view."
					switch selected {
					case "Dialogue":
						log.Print(logMessage)
						NavigateTo(window, directoryPath, MakeDialogueView)
					case "Installation":
						log.Print(logMessage)
						// NavigateTo(window, directoryPath, MakeInstallationView)
					case "Scripts":
						log.Print(logMessage)
						// NavigateTo(window, directoryPath, MakeScriptsView)
					default:
						log.Print("Failed to switch to different mod type")
						dialog.ShowError(fmt.Errorf("no modding available for %s", selected), window)
					}
				}
			}, window)
		}
		return
	}
}
