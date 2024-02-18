package views

import (
	"fyne.io/fyne/v2"
	"log"
)

// NavigateTo takes a fyne window and a destination view function and sets the window content to the new view.
func NavigateTo(window fyne.Window, directoryPath string, viewFunc func(string, fyne.Window) fyne.CanvasObject) {
	window.SetContent(viewFunc(directoryPath, window))
	log.Printf("Directory Path is: " + directoryPath + "\n")
}
