package views

import (
	"UI/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MakeHomeView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	toolbar := CreateToolbar(directoryPath, window)
	tree := utils.CreateFileTree(directoryPath, func(selected string) {
	})

	btnToDialogue := widget.NewButton("Dialogue Mod", func() {
		//NavigateTo(window, directoryPath, MakeDialogueView(directoryPath, window))
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
