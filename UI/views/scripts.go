package views

import (
	"UI/components"
	"UI/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"path/filepath"
)

func MakeScriptsView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Setting default variables
	newDirectoryPath := utils.GetScriptDirectory()
	defaultScriptsFilePath := filepath.Join(newDirectoryPath, "scripts_example.txt")
	templateScriptsFilePath := filepath.Join(newDirectoryPath, "scripts_template.txt")
	skeletonScriptsFilePath := filepath.Join(newDirectoryPath, "scripts_skeleton.txt")
	workingFilePath := filepath.Join(newDirectoryPath, "working.tmp")

	//Set Toolbar
	creatureID := ""
	modType := "Scripts"
	extension := ".baf"
	toolbar := CreateToolbar(directoryPath, window, creatureID, modType, extension, newDirectoryPath)

	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	contentLabel.Wrapping = fyne.TextWrapWord

	// Load and display the default file content
	fileContentView := utils.LoadFileContent(defaultScriptsFilePath)

	//Buttons for Initial Dialogue options
	btnToNextScriptsPage := widget.NewButton("Next", func() {
		NavigateTo(window, directoryPath, MakeScriptsModView)
	})

	// Create the file tree with double-click handling
	tree := utils.CreateFileTree(directoryPath, func(selected string) {
		// Single click actions, can go here.
		fullPath := filepath.Join(directoryPath, selected)
		content, err := ioutil.ReadFile(fullPath)
		// Handle the ReadFile error
		if err != nil {
			HandleErrorAndNavigate(err, newDirectoryPath, fullPath, selected, window)
		} else {
			// Show file content in a dialog
			fileContentDialog := dialog.NewCustom("File Content", "Close", widget.NewLabel(string(content)), window)
			fileContentDialog.Show()
		}
	}, func(selected string) {})

	// Hide Next button until New or Load is clicked.
	btnToNextScriptsPage.Hide()

	btnForNewScripts := widget.NewButton("New", func() {
		components.MakeNewFile(templateScriptsFilePath, newDirectoryPath, window)
		utils.UpdateFileContent(skeletonScriptsFilePath)
		tree.Refresh()
		btnToNextScriptsPage.Show()

	})

	btnForLoadModFile := components.CreateLoadModButton(window, ".d", newDirectoryPath, func() {
		utils.UpdateFileContent(workingFilePath)
		tree.Refresh()
		btnToNextScriptsPage.Show()
	})

	paddedButtonBar2 := container.NewHBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
		btnToNextScriptsPage,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	padButtonBar2 := container.NewVBox(
		layout.NewSpacer(),
		paddedButtonBar2,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	// VBox containing your buttons and file preview
	btnOptions := container.NewVBox(
		layout.NewSpacer(),
		btnForNewScripts,
		btnForLoadModFile,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	filePreview := container.NewVBox(
		layout.NewSpacer(),
		fileContentView,
		layout.NewSpacer(),
	)

	grid := container.NewAdaptiveGrid(4,
		layout.NewSpacer(),
		btnOptions,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	vBox := container.NewVBox(
		layout.NewSpacer(),
		grid,
		padButtonBar2,
	)

	vBox2 := container.NewVBox(
		toolbar,
		widget.NewLabel(""),
		vBox,
	)

	hBox := container.NewHBox(vBox2, filePreview)

	split := container.NewHSplit(tree, hBox)
	split.Offset = .15

	return split
}
