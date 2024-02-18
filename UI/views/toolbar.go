package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// CreateToolbar creates and returns a configured toolbar widget
func CreateToolbar(directoryPath string, window fyne.Window) *fyne.Container {
	homeButton := widget.NewButton("Home", func() {
		NavigateTo(window, directoryPath, MakeHomeView) //TODO: This needs to revert to original Directory Path.
	})
	exportButton := widget.NewButton("Export", func() {
		dialog.ShowInformation("Information", "WIP: Clicking this will display a confirmation to user.\n"+
			"Then it will trigger the backend to compile using the backend to be game ready.",
			window)
	})
	saveButton := widget.NewButton("Save", func() {
		dialog.ShowInformation("Saving", "WIP: Clicking this will allow you to Save the working Session File", window)
	})
	helpButton := widget.NewButton("Help", func() {
		dialog.ShowInformation("Help", "WIP: Clicking this will provide you a URL to our github to submit a ticket", window)
	})

	// Create the toolbar using a horizontal container
	toolbar := container.NewHBox(
		homeButton,
		exportButton,
		saveButton,
		helpButton,
	)

	return toolbar
}
