package views

import (
	"UI/components"
	"UI/utils"
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
	"strings"
)

func MakeNextInstallationView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Setting default variables
	newDirectoryPath := utils.GetInstallationDirectory()
	skeletonFilePath := filepath.Join(newDirectoryPath, "installation_skeleton.txt")
	var workingFilePath = filepath.Join(newDirectoryPath, "working.tmp") // Ensure this is mutable

	//Set Toolbar
	AuthorInput := ""
	modType := "Installation"
	extension := ".tp2"
	toolbar := CreateToolbar(directoryPath, window, AuthorInput, modType, extension, newDirectoryPath)

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
		// "AUTHORNAMEHERE" was not found in the file, fileContentView has been set
		displayFilePath = workingFilePath
		println("String 'AUTHORNAMEHERE' not found, fileContentView set to:", workingFilePath)
	} else {
		// "AUTHORNAMEHERE was found in the file
		displayFilePath = skeletonFilePath
		println("String 'AUTHORNAMEHERE' found, fileContentView set to:", skeletonFilePath)
	}
	fileContentView := utils.LoadFileContent(displayFilePath)

	inputAuthorBox := components.CreateLabeledTextInput("Author name:", func(inputValue string) {
		AuthorInput = inputValue
	})

	var VersionInput float64
	inputVersionBox := components.CreateLabeledTextInput("Version example: 1.0", func(versionStr string) {
		val, err := strconv.ParseFloat(versionStr, 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}
		VersionInput = val
	})

	var DialogueInput string
	btnSelectDialogueFile := components.SelectFilesForInstallation(window, ".d", utils.GetDialogueDirectory(), func(fileName string) {
		DialogueInput = strings.TrimSuffix(fileName, ".d")
		fmt.Println("The selected file is:", fileName)
	})

	var ScriptsInput string
	btnSelectScriptsFile := components.SelectFilesForInstallation(window, ".baf", utils.GetScriptDirectory(), func(fileName string) {
		ScriptsInput = strings.TrimSuffix(fileName, ".baf")
		fmt.Println("The selected file is:", fileName)
	})

	var CreatureFileInput string
	var CreatureNameInput string
	btnSelectCreatureFile := components.SelectFilesForInstallation(window, ".cre", utils.GetParentDirectory(), func(fileName string) {
		CreatureFileInput = strings.TrimSuffix(fileName, ".cre")
		fmt.Println("The selected file is:", fileName)
		CreatureNameInput = CreatureFileInput
	})

	variableBoxesContainer := container.NewVBox(
		inputAuthorBox,
		inputVersionBox,
		btnSelectDialogueFile,
		btnSelectScriptsFile,
		btnSelectCreatureFile,
	)

	btnToSave := widget.NewButton("Save", func() {
		if AuthorInput == "" {
			// Handle the case where AuthorInput is not set
			dialog.ShowInformation("Error", "Please enter an Author name", window)
			return
		}
		components.SaveFile(AuthorInput, "installation", extension, newDirectoryPath, window)
		NavigateTo(window, directoryPath, MakeHomeView)
	})
	btnToSave.Hide()

	insertButton := widget.NewButton("Insert", func() {
		userInputs := &components.UserInputs{
			Author:       AuthorInput,
			Version:      VersionInput,
			ModName:      utils.GetModFolder(),
			DialogueFile: DialogueInput,
			ScriptsFile:  ScriptsInput,
			CreatureFile: CreatureFileInput,
			CreatureName: CreatureNameInput,
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
		layout.NewSpacer(),
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
