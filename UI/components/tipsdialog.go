package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowTipsDialog(window fyne.Window) {
	// Define the text blurbs
	tips := []string{
		"Changelog Example:\n0.0.1-beta-01 Release\n- Added dialogs for each of the toolbar buttons\n" +
			"- Full file modification for Dialogue\n" +
			"- This is just an Example Demonstration",
		"Tip 1:\nThis will be a place for us to provide users\nsome tips and tricks\nWhich will be based on user feedback\nand any changes from the release",
		"Tip 2:\nAnother tip...\nThis a Demo and is meant to be fun\n\n" +
			"Knock, Knock\nWho's there?\nGit\nGit who?\nGit commit so we can track who made this knock-knock joke!",
	}

	// Current index of the tip
	var currentTip int
	textLabel := widget.NewLabel(tips[currentTip])
	textLabel.Alignment = fyne.TextAlignCenter

	// Function to update the text label
	updateTextLabel := func() {
		textLabel.SetText(tips[currentTip])
	}

	// Next and Previous buttons
	nextBtn := widget.NewButton("Next", func() {
		if currentTip < len(tips)-1 {
			currentTip++
			updateTextLabel()
		}
	})

	prevBtn := widget.NewButton("Prev", func() {
		if currentTip > 0 {
			currentTip--
			updateTextLabel()
		}
	})

	// Button layout with spacers for horizontal centering
	buttonLayout := container.NewHBox(layout.NewSpacer(), prevBtn, nextBtn, layout.NewSpacer())

	// Content layout with vertical alignment
	tipsContent := container.NewVBox(
		textLabel,
	)

	image := canvas.NewImageFromFile("src/image-100x100.jpg")
	image.FillMode = canvas.ImageFillOriginal

	imageContain := container.NewHBox(image, layout.NewSpacer())

	content := container.NewHBox(
		imageContain,
		layout.NewSpacer(),
		tipsContent,
	)

	contain := container.NewVBox(
		content,
		buttonLayout,
	)

	// Custom dialog
	customDialog := dialog.NewCustom("Tips and Tricks", "OK", contain, window)
	customDialog.SetOnClosed(func() {})

	customDialog.Show()
}
