package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"goTesting/components"
	"goTesting/utils"
)

var selectedDirectoryPath string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("File Tree Viewer")

	content := widget.NewLabel("Select a directory...")

	// Button to allow user to select a directory
	openProjectBtn := widget.NewButton("Select Directory", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil || err != nil {
				return
			}
			selectedDirectoryPath = uri.Path() // Store the selected directory path globally
			tree := utils.CreateFileTree(selectedDirectoryPath, func(selected string) {
				content.SetText("Selected: " + selected)
			})
			myWindow.SetContent(container.NewHSplit(tree, content)) // Update the window content with the new tree
		}, myWindow)
	})

	newProjectBtn := widget.NewButton("New Project", func() {
		utils.PromptForProjectName(myWindow, func(newPath string) {
			selectedDirectoryPath = newPath // Update the global variable with the new path
			templatePath := "dialogue_temp.txt"
			components.MakeNewFile(templatePath, selectedDirectoryPath, myWindow)

			tree := utils.CreateFileTree(selectedDirectoryPath, func(selected string) {
				content.SetText("Selected: " + selected)
			})
			myWindow.SetContent(container.NewHSplit(tree, content)) // Update the window content with the new tree
		})
	})

	// Initial content with the select directory button
	initialContent := container.NewVBox(
		openProjectBtn,
		newProjectBtn,
		content,
	)

	myWindow.SetContent(initialContent)
	myWindow.Resize(fyne.NewSize(1280, 800))
	myWindow.ShowAndRun()
}
