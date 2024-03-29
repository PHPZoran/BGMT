package views

import (
	"UI/components"
	"UI/utils"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
)

func MakeSpeakerView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Setting default variables
	newDirectoryPath := utils.GetDialogueDirectory()
	skeletonFilePath := filepath.Join(newDirectoryPath, "dialogue_skeleton.txt")
	var workingFilePath = filepath.Join(newDirectoryPath, "working.tmp") // Ensure this is mutable

	//Set Toolbar
	creatureID := ""
	modType := "Dialogue"
	var extension = ".d"
	toolbar := CreateToolbar(directoryPath, window, creatureID, modType, extension, newDirectoryPath)

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

	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	contentLabel.Wrapping = fyne.TextWrapWord

	// Load and display the default file content
	displayFilePath := ""
	found := utils.CheckFileForString(workingFilePath)
	if !found {
		// "$creatureID" was not found in the file, fileContentView has been set
		displayFilePath = workingFilePath
		println("String 'CREATUREID' not found, fileContentView set to:", workingFilePath)
	} else {
		// "$creatureID" was found in the file
		displayFilePath = skeletonFilePath
		println("String 'CREATUREID' found, fileContentView set to:", skeletonFilePath)
	}
	fileContentView := utils.LoadFileContent(displayFilePath)

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

	var ValueInput int64
	inputValueBox := components.CreateLabeledTextInput("Value: numb -255> and <255", func(inputValue string) {
		val, err := strconv.ParseInt(inputValue, 10, 32)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		ValueInput = val
	})

	var VariableInput2 string
	inputVariable2Box := components.CreateLabeledTextInput("Variable 2 example: KING123", func(inputValue string) {
		VariableInput2 = inputValue
	})

	var TypeInput2 string
	inputType2Box := components.CreateLabeledTextInput("Type: GLOBAL or LOCALS", func(inputValue string) {
		TypeInput2 = inputValue
	})

	var ValueInput2 int64
	inputValue2Box := components.CreateLabeledTextInput("Value: numb -255> and <255", func(inputValue string) {
		val, err := strconv.ParseInt(inputValue, 10, 32)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		ValueInput2 = val
	})

	variableBoxesContainer := container.NewVBox()

	triggersChoiceBox := widget.NewSelect([]string{"Template 1"}, func(value string) {
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

	btnToSave := widget.NewButton("Save", func() {
		if CreatureID == "" {
			// Handle the case where creatureID is not set
			dialog.ShowInformation("Error", "Please enter a Speaker ID", window)
			return
		}
		components.SaveFile(CreatureID, "dialogue", extension, newDirectoryPath, window)
		NavigateTo(window, directoryPath, MakeHomeView)
	})
	btnToSave.Hide()

	insertButton := widget.NewButton("Insert", func() {
		err := validateDialogueInputs(TypeInput, ValueInput, TypeInput2, ValueInput2)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

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

	btnFmtHBox := container.NewAdaptiveGrid(9,
		layout.NewSpacer(),
		insertButton,
		btnToSave,
		layout.NewSpacer(),
	)

	paddedInfoBtn := container.NewHBox(
		layout.NewSpacer(),
		infoButton,
		layout.NewSpacer(),
	)

	paddedBtnBox := container.NewVBox(
		layout.NewSpacer(),
		btnFmtHBox,
		layout.NewSpacer(),
		layout.NewSpacer(),
		paddedInfoBtn,
	)

	userInputFmtBox := container.NewVBox(
		layout.NewSpacer(),
		inputSpeakerIDBox,
		triggersChoiceBox,
		variableBoxesContainer,
		layout.NewSpacer(),
		paddedBtnBox,
		layout.NewSpacer(),
	)

	contentViewFmtBox := container.NewVBox(
		layout.NewSpacer(),
		contentLabel,
		fileContentView,
	)

	mainContentBox := container.NewHBox(
		userInputFmtBox,
		layout.NewSpacer(),
		contentViewFmtBox,
	)

	vbox := container.NewVBox(
		toolbar,
		mainContentBox,
		layout.NewSpacer(),
	)
	split := container.NewHSplit(tree, vbox)
	split.Offset = .15
	return container.NewVBox(split)
}

func validateDialogueInputs(typeInput string, valueInput int64, typeInput2 string, valueInput2 int64) error {
	if !isValidType(typeInput) {
		return errors.New("please enter a valid type: LOCALS or GLOBAL")
	}
	if !isValidValue(valueInput, -255, 255) {
		return errors.New("please enter a valid value between -255 and 255")
	}
	if !isValidType(typeInput2) {
		return errors.New("please enter a valid type: LOCALS or GLOBAL")
	}
	if !isValidValue(valueInput2, -255, 255) {
		return errors.New("please enter a valid value between -255 and 255")
	}
	return nil
}
