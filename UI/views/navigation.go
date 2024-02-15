package views

import (
	"fyne.io/fyne/v2"
)

// NavigateTo takes a fyne window and a destination view function and sets the window content to the new view.
func NavigateTo(window fyne.Window, viewFunc func(fyne.Window) fyne.CanvasObject) {
	window.SetContent(viewFunc(window))
}
