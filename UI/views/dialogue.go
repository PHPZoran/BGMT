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

func MakeDialogueView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Setting default variables
	newDirectoryPath := utils.GetDialogueDirectory()
	defaultDialogueFilePath := filepath.Join(newDirectoryPath, "dialogue_example.txt")
	templateDialogueFilePath := filepath.Join(newDirectoryPath, "dialogue_temp.txt")
	skeletonDialogueFilePath := filepath.Join(newDirectoryPath, "dialogue_skeleton.txt")
	workingFilePath := filepath.Join(newDirectoryPath, "working.tmp")

	//Set Toolbar
	creatureID := ""
	modType := "Dialogue"
	extension := ".d"
	toolbar := CreateToolbar(directoryPath, window, creatureID, modType, extension, newDirectoryPath)

	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	contentLabel.Wrapping = fyne.TextWrapWord

	// Load and display the default file content
	fileContentView := utils.LoadFileContent(defaultDialogueFilePath)

	//Buttons for Initial Dialogue options
	btnToNextDialoguePage := widget.NewButton("Next", func() {
		NavigateTo(window, directoryPath, MakeSpeakerView)
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
	btnToNextDialoguePage.Hide()

	btnForNewDialogue := widget.NewButton("New", func() {
		components.MakeNewFile(templateDialogueFilePath, newDirectoryPath, window)
		utils.UpdateFileContent(skeletonDialogueFilePath)
		tree.Refresh()
		btnToNextDialoguePage.Show()

	})

	btnForLoadModFile := components.CreateLoadModButton(window, ".d", newDirectoryPath, func() {
		utils.UpdateFileContent(workingFilePath)
		tree.Refresh()
		btnToNextDialoguePage.Show()
	})

	paddedButtonBar2 := container.NewHBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
		btnToNextDialoguePage,
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
		btnForNewDialogue,
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
