package views

import (
	"UI/components"
	"UI/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MakeDialogueView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	directoryPath += "/Dialogue"
	toolbar := CreateToolbar(directoryPath, window)
	tree := utils.CreateFileTree(directoryPath, func(selected string) {
	})
	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	contentLabel.Wrapping = fyne.TextWrapWord
	fileContentView := utils.LoadFileContent(directoryPath + "/dialogue_example.txt")

	//Buttons for Initial Dialogue options
	btnToNextDialoguePage := widget.NewButton("Next", func() {
		NavigateTo(window, directoryPath, MakeSpeakerView)
	})

	// Hide Next button until New or Load is clicked.
	btnToNextDialoguePage.Hide()

	btnForNewDialogue := widget.NewButton("New", func() {
		components.MakeNewFile(directoryPath+"/dialogue_temp.txt", directoryPath, window)
		utils.UpdateFileContent(directoryPath + "/dialogue_skeleton.txt")
		btnToNextDialoguePage.Show()
	})

	btnForLoadModFile := components.CreateLoadModButton(window, ".d", func() {
		utils.UpdateFileContent(directoryPath + "/working.tmp")
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

//DONE: Generate a true skeleton template of dialogue as dialogue_skeleton.txt
//DONE: Make sure NEWBTN creates the "working.tmp". Should be a copy of dialogue_skeleton.txt
//DONE: Display dialogue_example.txt with real values initially
//DONE: Refactor naming of speakerID to creatureID.
//DONE: generate dialogueID using random 6 char and numbers.
//DONE: Set parameters for creatureID and variable, numerical parameters for value, "LOCALS" or "GLOBAL" options for type.
//TODO: Help Button on speaker.go
//DONE: Change triggerChoices to then display more options on current window (not a pop-up)
