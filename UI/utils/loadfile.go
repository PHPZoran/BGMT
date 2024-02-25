package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
	"strings"
)

var contentLabel *widget.Label

// LoadFileContent updates the label's text with the content of the specified file.
func LoadFileContent(filename string) *fyne.Container {
	contentLabel = widget.NewLabel("")
	UpdateFileContent(filename) // Load initial content

	// Create a padded container for the label
	paddedContainer := NewPaddedLabel(contentLabel)
	return paddedContainer
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

// UpdateFileContent Function to load file content and update the label's text
func UpdateFileContent(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read file:", err)
		contentLabel.SetText("Failed to load content")
		return
	}
	contentLabel.SetText(string(content))
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

// NewPaddedLabel creates a new label with specified padding, background, and text color.
func NewPaddedLabel(label *widget.Label) *fyne.Container {
	// Create a background rectangle with the desired color
	background := canvas.NewRectangle(theme.InputBackgroundColor())

	// Wrap the label in a scroll container to allow scrolling if the content overflows.
	scrollContainer := container.NewScroll(container.NewStack(background, label))
	scrollContainer.SetMinSize(fyne.NewSize(400, 500)) // Ensure the container is large enough to display content

	// Return a container with the specified padding around the label
	return container.NewPadded(scrollContainer)
}

func CheckFileForString(workingFilePath string) bool {
	// Read the content of the file
	content, err := ioutil.ReadFile(workingFilePath)
	if err != nil {
		// Handle the error according to your application's requirements
		panic(err) // Example error handling
	}

	// Convert the content to a string
	contentStr := string(content)

	// Check for "$creatureID" or "AUTHORNAMEHERE"
	if strings.Contains(contentStr, "CREATUREID") || strings.Contains(contentStr, "AUTHORNAMEHERE") {
		return true
	}
	return false
}
