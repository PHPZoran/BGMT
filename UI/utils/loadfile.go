package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
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
	//label := widget.NewLabel(data)
	//label.Wrapping = fyne.TextWrapWord // Enable word wrapping

	// Create a background rectangle with the desired color
	background := canvas.NewRectangle(theme.InputBackgroundColor())
	//background.FillColor = color.White

	// Wrap the label in a scroll container to allow scrolling if the content overflows.
	scrollContainer := container.NewScroll(container.NewStack(background, label))
	scrollContainer.SetMinSize(fyne.NewSize(400, 500)) // Ensure the container is large enough to display content

	// Return a container with the specified padding around the label
	return container.NewPadded(scrollContainer)
}
