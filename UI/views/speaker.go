package views

import (
	"UI/components"
	"UI/utils"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"path/filepath"
)

func MakeSpeakerView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Set Toolbar --- Remove?
	toolbar := CreateToolbar(directoryPath, window)

	//Setting default variables
	newDirectoryPath := filepath.Join(directoryPath, "Dialogue")
	skeletonFilePath := filepath.Join(newDirectoryPath, "dialogue_skeleton.txt")
	var workingFilePath = filepath.Join(newDirectoryPath, "working.tmp") // Ensure this is mutable

	// Create the file tree with double-click handling
	tree := utils.CreateFileTree(newDirectoryPath, func(selected string) {
		// Single click actions, can go here.
		fullPath := filepath.Join(newDirectoryPath, selected)
		utils.UpdateFileContent(fullPath)
		// Confirmation dialog
		dialog.ShowConfirm("Change File Confirmation", "Are you sure? Any changes will be lost?",
			func(confirm bool) {
				if confirm {
					// If user confirms, update workingFilePath and refresh content view
					workingFilePath = fullPath
					utils.UpdateFileContent(workingFilePath)
				}
				utils.UpdateFileContent(skeletonFilePath)
			}, window)
	}, func(selected string) {
		// Handle double-click on a file
	})

	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	contentLabel.Wrapping = fyne.TextWrapWord

	// Load and display the default file content
	fileContentView := utils.LoadFileContent(skeletonFilePath)

	var CreatureID string
	inputSpeakerIDBox := components.CreateLabeledTextInput("CreatureID example: KINGKONG", func(inputValue string) {
		CreatureID = inputValue
	})

	var VariableInput string
	inputVariableBox := components.CreateLabeledTextInput("Variable example: KONG123", func(inputValue string) {
		VariableInput = inputValue
	})

	var TypeInput string
	inputTypeBox := components.CreateLabeledTextInput("Type: GLOBAL or LOCALS", func(inputValue string) {
		TypeInput = inputValue
	})

	var ValueInput string
	inputValueBox := components.CreateLabeledTextInputInt(window, "Value: numb -255> and <255", func(inputValue string) {
		ValueInput = inputValue
	})

	var VariableInput2 string
	inputVariable2Box := components.CreateLabeledTextInput("Variable 2 example: KING123", func(inputValue string) {
		VariableInput2 = inputValue
	})

	var TypeInput2 string
	inputType2Box := components.CreateLabeledTextInput("Type: GLOBAL or LOCALS", func(inputValue string) {
		TypeInput2 = inputValue
	})

	var ValueInput2 string
	inputValue2Box := components.CreateLabeledTextInputInt(window, "Value: numb -255> and <255", func(inputValue string) {
		ValueInput2 = inputValue
	})

	variableBoxesContainer := container.NewVBox()

	triggersChoiceBox := widget.NewSelect([]string{"Template 1", "Template 2", "Template 3"}, func(value string) {
		switch value {
		case "Template 1", "Template 2", "Template 3":
			variableBoxesContainer.Add(inputVariableBox)
			variableBoxesContainer.Add(inputTypeBox)
			variableBoxesContainer.Add(inputValueBox)
			variableBoxesContainer.Add(inputVariable2Box)
			variableBoxesContainer.Add(inputType2Box)
			variableBoxesContainer.Add(inputValue2Box)
		}
		variableBoxesContainer.Refresh()
	})

	// VBox containing your button and label
	vbox1 := container.NewVBox(
		layout.NewSpacer(),
		contentLabel,
		fileContentView,
	)

	vbox2 := container.NewVBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
		inputSpeakerIDBox,
		triggersChoiceBox,
		variableBoxesContainer,
		layout.NewSpacer(),
	)

	grid := container.NewAdaptiveGrid(3,
		vbox2,
		layout.NewSpacer(),
		vbox1,
	)

	var extension = ".d"
	btnToSave := widget.NewButton("Save", func() {
		if CreatureID == "" {
			// Handle the case where speakerID is not set
			dialog.ShowInformation("Error", "Please enter a Speaker ID", window)
			return
		}
		components.SaveFile(CreatureID, "dialogue", extension, window)
		NavigateTo(window, directoryPath, MakeHomeView)
	})
	btnToSave.Hide()

	insertButton := widget.NewButton("Insert", func() {
		userInputs := &components.UserInputs{
			Variable:   VariableInput,
			Type:       TypeInput,
			Value:      ValueInput,
			Variable2:  VariableInput2,
			Type2:      TypeInput2,
			Value2:     ValueInput2,
			DialogueID: utils.GenerateRandomString(),
			CreatureID: CreatureID,
		}
		log.Printf("Inserting UserInputs: %+v\n", userInputs)
		// Insert template UserInputs into temp file
		templateErr := components.InsertUserInputs(workingFilePath, userInputs)
		if templateErr != nil {
			return
		}
		utils.UpdateFileContent(workingFilePath)
		btnToSave.Show()
	})

	infoButton := widget.NewButtonWithIcon("Info", theme.InfoIcon(), func() {
		dialog.ShowInformation("Information", "CreatureID is the id of the character you are working on.\n"+
			"Variable is unique to this dialogue.\n"+
			"Value determines if its local or global.",
			window)
	})

	paddedToolbar2 := container.NewHBox(
		layout.NewSpacer(),
		insertButton,
		btnToSave,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	paddedToolbar3 := container.NewHBox(
		layout.NewSpacer(),
		infoButton,
		layout.NewSpacer(),
	)

	padToolbar2 := container.NewVBox(
		layout.NewSpacer(),
		paddedToolbar2,
		layout.NewSpacer(),
		layout.NewSpacer(),
		paddedToolbar3,
	)

	vbox := container.NewVBox(
		toolbar,
		grid,
		padToolbar2,
	)
	split := container.NewHSplit(tree, vbox)
	split.Offset = .3
	//TODO: FIX THE LAYOUT SO TREE IS VIEWABLE
	return container.NewVBox(split)
}
