package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CreateLabeledTextInput creates and returns a new container with a label and a text input widget
func CreateLabeledTextInput(placeholder string, onChange func(string)) fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetPlaceHolder(placeholder)
	input.OnChanged = onChange

	// Combine the label and text input in a vertical box layout
	return container.NewVBox(
		input,
	)
}

func ShowFileLoadDialog(window fyne.Window, expectedExtension string, onValidFileSelected func(filePath string), initialDir string) {
	// Convert initial directory path to ListableURI
	initialDirURI, err := storage.ListerForURI(storage.NewFileURI(initialDir))
	if err != nil {
		dialog.ShowError(err, window)
		return
	}

	// Create a file open dialog with the custom callback
	fileDialog := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		if file == nil {
			return
		}
		// Check if the file has the appropriate extension
		if filepath.Ext(file.URI().Path()) == expectedExtension {
			// Execute the callback with the selected file path
			onValidFileSelected(file.URI().Path())
		} else {
			dialog.ShowInformation("Invalid File", "Please select a file with a "+expectedExtension+" extension", window)
		}
	}, window)
	// Set the initial location for the dialog
	fileDialog.SetLocation(initialDirURI)
	fileDialog.Show()
}

// CreateLoadModButton creates and returns a button for loading and copying a mod file.
func CreateLoadModButton(window fyne.Window, expectedExtension string, initialDir string, postProcess func()) *widget.Button {
	btn := widget.NewButton("Load Mod File", func() {
		ShowFileLoadDialog(window, expectedExtension, func(srcFilePath string) {
			// File is valid, now copy to working.tmp
			destPath := filepath.Join(initialDir, "working.tmp")
			fmt.Println("Copying from:", srcFilePath, "to:", destPath)
			err := CopyFile(srcFilePath, destPath)
			if err != nil {
				dialog.ShowError(err, window)
			} else {
				dialog.ShowInformation("Success", "File loaded and copied to working.tmp", window)
				if postProcess != nil {
					postProcess() // Execute any post-processing function if provided
				}
			}
		}, initialDir)
	})

	return btn
}

func SelectFilesForInstallation(window fyne.Window, expectedExtension string, initialDir string, onFileSelected func(fileName string)) *widget.Button {
	var inputFileName string
	btn := widget.NewButton("Select Mod File", func() {
		ShowFileLoadDialog(window, expectedExtension, func(srcFilePath string) {
			// File is valid, now copy to working.tmp
			switch expectedExtension {
			case ".d", ".baf", ".cre":
				inputFileName = filepath.Base(srcFilePath)
				fmt.Println("File chosen is:" + inputFileName)
			}
			fmt.Println("File chosen is:" + inputFileName)
			onFileSelected(inputFileName)
		}, initialDir)
	})
	return btn
}

// CopyFile copies the contents of the src file to the dst file.
func CopyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, input, 0644)
}

func MakeNewFile(templatePath, projectDir string, window fyne.Window) {
	dst := filepath.Join(projectDir, "working.tmp")
	copyErr := CopyFile(templatePath, dst)
	if copyErr != nil {
		dialog.ShowError(copyErr, window)
		return
	}
	dialog.ShowInformation("Success", "New project file created", window)
}

// InsertUserInputs replaces placeholders in a file with user inputs.
func InsertUserInputs(filePath string, inputs *UserInputs) error {

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("ERROR: Something is wrong")
		return err
	}

	// Convert the content to a string for replacement
	newContent := string(content)

	// Replace each placeholder with the corresponding value from UserInputs, if not empty
	if inputs.Variable != "" {
		newContent = strings.ReplaceAll(newContent, "VARIABLENAME", inputs.Variable)
	}
	if inputs.Type != "" {
		newContent = strings.ReplaceAll(newContent, "TYPE", inputs.Type)
	}
	valueStr := fmt.Sprintf("%v", inputs.Value)
	newContent = strings.ReplaceAll(newContent, "INT", valueStr)

	if inputs.Variable2 != "" {
		newContent = strings.ReplaceAll(newContent, "VAR", inputs.Variable2)
	}
	if inputs.Type2 != "" {
		newContent = strings.ReplaceAll(newContent, "TYP", inputs.Type2)
	}
	valueStr2 := fmt.Sprintf("%v", inputs.Value2)
	newContent = strings.ReplaceAll(newContent, "VAL", valueStr2)

	if inputs.CreatureID != "" {
		newContent = strings.ReplaceAll(newContent, "CREATUREID", inputs.CreatureID)
	}
	if inputs.DialogueID != "" {
		newContent = strings.ReplaceAll(newContent, "$dialogueID", inputs.DialogueID)
	}
	if inputs.Author != "" {
		newContent = strings.ReplaceAll(newContent, "AUTHORNAMEHERE", inputs.Author)
	}
	if inputs.ModName != "" {
		newContent = strings.ReplaceAll(newContent, "MODNAMEHERE", inputs.ModName)
	}
	if inputs.Version != 0 {
		versionStr := fmt.Sprintf("%v", inputs.Version)
		newContent = strings.ReplaceAll(newContent, "VERSIONNUMHERE", versionStr)
	}
	if inputs.DialogueFile != "" {
		newContent = strings.ReplaceAll(newContent, "DIALOGUENAMEHERE", inputs.DialogueFile)
	}
	if inputs.ScriptsFile != "" {
		newContent = strings.ReplaceAll(newContent, "SCRIPTNAMEHERE", inputs.ScriptsFile)
	}
	if inputs.CreatureFile != "" {
		newContent = strings.ReplaceAll(newContent, "CREATUREFILENAMEHERE", inputs.CreatureFile)
	}
	if inputs.CreatureName != "" {
		newContent = strings.ReplaceAll(newContent, "CREATUREDISPLAYNAMEHERE", inputs.CreatureName)
	}
	xValueStr := fmt.Sprintf("%v", inputs.XCoordinate)
	newContent = strings.ReplaceAll(newContent, "XCOORDINATE", xValueStr)

	yValueStr := fmt.Sprintf("%v", inputs.YCoordinate)
	newContent = strings.ReplaceAll(newContent, "YCOORDINATE", yValueStr)

	if inputs.FacingDirect != "" {
		newContent = strings.ReplaceAll(newContent, "FACINGDIRECTION", inputs.FacingDirect)
	}

	// Write the updated content back to the file
	if err := os.WriteFile(filePath, []byte(newContent), 0644); err != nil {
		log.Printf("Failed to write to file '%s': %v", filePath, err)
		return err
	}
	log.Printf("Successfully wrote to file '%s'", filePath)
	// After writing to the file
	readBackContent, _ := os.ReadFile(filePath)
	log.Printf("Content after write: %s", string(readBackContent))

	return nil
}

type UserInputs struct {
	Variable     string
	Type         string
	Value        int64
	Variable2    string
	Type2        string
	Value2       int64
	DialogueID   string
	CreatureID   string
	Author       string
	ModName      string
	DialogueFile string
	ScriptsFile  string
	CreatureFile string
	CreatureName string
	FacingDirect string
	XCoordinate  int64
	YCoordinate  int64
	Version      float64
}
