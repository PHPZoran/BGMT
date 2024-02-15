package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"goTesting/components"
	"goTesting/utils"
)

func MakeDialogueView(window fyne.Window) fyne.CanvasObject {
	toolbar := CreateToolbar(window)

	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	utils.LoadFileContent("dialogue_example.txt", contentLabel)
	contentLabel.Wrapping = fyne.TextWrapWord

	//Buttons for Initial Dialogue options
	btnToNextDialoguePage := widget.NewButton("Next", func() {
		NavigateTo(window, MakeSpeakerView)
	})
	// Hide Next button until New or Load is clicked.
	btnToNextDialoguePage.Hide()

	btnForNewDialogue := widget.NewButton("New", func() {
		components.MakeNewFile(window)
		utils.LoadFileContent("dialogue_skeleton.txt", contentLabel)
		btnToNextDialoguePage.Show()
	})

	btnForLoadModFile := components.CreateLoadModButton(window, ".d", func() {
		utils.LoadFileContent("working.tmp", contentLabel)
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
		contentLabel,
	)

	grid := container.NewAdaptiveGrid(3,
		btnOptions,
		layout.NewSpacer(),
		filePreview,
	)

	vbox := container.NewVBox(
		toolbar,
		grid,
		padButtonBar2,
	)

	return container.NewVBox(vbox)
}

//DONE: Generate a true skeleton template of dialogue as dialogue_skeleton.txt
//DONE: Make sure NEWBTN creates the "working.tmp". Should be a copy of dialogue_skeleton.txt
//DONE: Display dialogue_example.txt with real values initially
//DONE: Refactor naming of speakerID to creatureID.
//DONE: generate dialogueID using random 6 char and numbers.
//DONE: Set parameters for creatureID and variable, numerical parameters for value, "LOCALS" or "GLOBAL" options for type.
//TODO: Help Button on speaker.go
//DONE: Change triggerChoices to then display more options on current window (not a pop-up)
