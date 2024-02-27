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

func MakeScriptsModView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Setting default variables
	newDirectoryPath := utils.GetScriptDirectory()
	skeletonFilePath := filepath.Join(newDirectoryPath, "scripts_skeleton.txt")
	var workingFilePath = filepath.Join(newDirectoryPath, "working.tmp") // Ensure this is mutable

	//Set Toolbar
	creatureID := ""
	modType := "Scripts"
	var extension = ".baf"
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
		// "CREATUREID" was not found in the file, fileContentView has been set
		displayFilePath = workingFilePath
		println("String 'CREATUREID' not found, fileContentView set to:", workingFilePath)
	} else {
		// "CREATUREID" was found in the file
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

	var FacingDirectInput string
	inputDirectionBox := components.CreateLabeledTextInput("Facing Directions example: N for North", func(inputValue string) {
		FacingDirectInput = inputValue
	})

	var XCoordinateInput int64
	inputXCoordBox := components.CreateLabeledTextInput("X Coordinate of Creature: 0 - 9999", func(inputValue string) {
		val, err := strconv.ParseInt(inputValue, 10, 32)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		XCoordinateInput = val
	})

	var YCoordinateInput int64
	inputYCoordBox := components.CreateLabeledTextInput("Y Coordinate of Creature: 0 - 9999", func(inputValue string) {
		val, err := strconv.ParseInt(inputValue, 10, 32)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		YCoordinateInput = val
	})

	variableBoxesContainer := container.NewVBox()

	triggersChoiceBox := widget.NewSelect([]string{"Template 1"}, func(value string) {
		switch value {
		case "Template 1":
			variableBoxesContainer.Add(inputVariableBox)
			variableBoxesContainer.Add(inputTypeBox)
			variableBoxesContainer.Add(inputValueBox)
			variableBoxesContainer.Add(inputDirectionBox)
			variableBoxesContainer.Add(inputXCoordBox)
			variableBoxesContainer.Add(inputYCoordBox)
		}
		variableBoxesContainer.Refresh()
	})

	btnToSave := widget.NewButton("Save", func() {
		if CreatureID == "" {
			// Handle the case where creatureID is not set
			dialog.ShowInformation("Error", "Please enter a CreatureID", window)
			return
		}
		components.SaveFile(CreatureID, modType, extension, newDirectoryPath, window)
		NavigateTo(window, directoryPath, MakeHomeView)
	})
	btnToSave.Hide()

	insertButton := widget.NewButton("Insert", func() {
		err := validateScriptsInputs(FacingDirectInput, TypeInput, ValueInput, XCoordinateInput, YCoordinateInput)
		if err != nil {
			dialog.ShowError(err, window)
			return
		}

		userInputs := &components.UserInputs{
			Variable:     VariableInput,
			Type:         TypeInput,
			Value:        ValueInput,
			FacingDirect: FacingDirectInput,
			XCoordinate:  XCoordinateInput,
			YCoordinate:  YCoordinateInput,
			CreatureID:   CreatureID,
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

// isValidDirection checks if the given direction is one of the valid directions.
func isValidDirection(direction string) bool {
	validDirections := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	for _, validDirection := range validDirections {
		if direction == validDirection {
			return true
		}
	}
	return false
}

// isValidDirection checks if the given direction is one of the valid directions.
func isValidType(typeInput string) bool {
	validTypes := []string{"LOCALS", "GLOBAL"}
	for _, validDirection := range validTypes {
		if typeInput == validDirection {
			return true
		}
	}
	return false
}

// isValidValue checks if the given value is in the appropriate range.
func isValidValue(value int64, min int64, max int64) bool {
	return value >= min && value <= max
}

func validateScriptsInputs(facingDirectInput string, typeInput string, valueInput int64, xCoordinateInput int64, yCoordinateInput int64) error {
	if !isValidDirection(facingDirectInput) {
		return errors.New("please enter a valid direction: N, NE, E, SE, S, SW, W, NW")
	}
	if !isValidType(typeInput) {
		return errors.New("please enter a valid type: LOCALS or GLOBAL")
	}
	if !isValidValue(valueInput, -255, 255) {
		return errors.New("please enter a valid value between -255 and 255")
	}
	if !isValidValue(xCoordinateInput, 0, 9999) {
		return errors.New("please enter a valid X coordinate between 0 and 9999")
	}
	if !isValidValue(yCoordinateInput, 0, 9999) {
		return errors.New("please enter a valid X coordinate between 0 and 9999")
	}
	return nil
}
