package views

import (
	"UI/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"path/filepath"
)

func MakeHomeView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Set Toolbar
	speakerID := ""
	modType := ""
	var extension = ""
	toolbar := CreateToolbar(directoryPath, window, speakerID, modType, extension, "")

	// Create the file tree with double-click handling
	tree := utils.CreateFileTree(directoryPath, func(selected string) {
		// Single click actions, can go here.
		fullPath := filepath.Join(directoryPath, selected)
		content, err := ioutil.ReadFile(fullPath)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		// Show file content in a dialog
		fileContentDialog := dialog.NewCustom("File Content", "Close", widget.NewLabel(string(content)), window)
		fileContentDialog.Show()
	}, func(selected string) {
		// Handle double-click on a file
	})

	btnToDialogue := widget.NewButton("Dialogue Mod", func() {
		NavigateTo(window, directoryPath, MakeDialogueView)
	})
	btnToScripts := widget.NewButton("Scripts Mod", func() {
		//utils.NavigateToScripts(window)
	})
	btnToModInstallation := widget.NewButton("Installation Mod", func() {
		//utils.NavigateToScripts(window)
	})

	vbox := container.NewVBox(
		btnToDialogue,
		layout.NewSpacer(),
		btnToScripts,
		layout.NewSpacer(),
		btnToModInstallation,
	)

	grid := container.NewGridWithColumns(5,
		layout.NewSpacer(),
		layout.NewSpacer(),
		vbox,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	vboxBtnLayout := container.NewGridWithRows(5,
		layout.NewSpacer(),
		layout.NewSpacer(),
		grid,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	split := container.NewHSplit(tree, vboxBtnLayout)
	split.Offset = .15

	return container.NewVBox(
		toolbar,
		widget.NewLabel(""),
		split,
	)
}
