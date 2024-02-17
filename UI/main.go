package main

import (
	"UI/utils"
	"UI/views"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
)

type AppState struct {
	SelectedDirectoryPath string
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Baldur's Gate Mod Tool")
	state := &AppState{}
	content := widget.NewLabel("Select a directory...")

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

	//File -> Menu Options
	menuItemExportProject := fyne.NewMenuItem("Export Project", nil)
	menuItemNewProject := fyne.NewMenuItem("New Project", func() {
		utils.PromptForProjectName(myWindow, func(newPath string) {
			state.SelectedDirectoryPath = newPath // Update the global variable with the new path

			homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
			myWindow.SetContent(homeView) // Update the window content with the new tree
		})
		menuItemExportProject.Disabled = false
	})
	menuItemOpenProject := fyne.NewMenuItem("Open Project", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil || err != nil {
				return
			}
			state.SelectedDirectoryPath = uri.Path() // Store the selected directory path globally

			if _, err := os.Stat(state.SelectedDirectoryPath + "/BGM.ini"); err != nil {
				dialog.ShowError(errors.New("the selected folder is not a BGM Project folder"), myWindow)
				return
			}

			homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
			myWindow.SetContent(homeView)
		}, myWindow)
		menuItemExportProject.Disabled = false
	})

	menuItemExportProject.Disabled = true

	//Help -> Menu Options
	menuItemReport := fyne.NewMenuItem("Report Bug", nil)
	menuItemSettings := fyne.NewMenuItem("Settings", nil)
	menuItemHelp := fyne.NewMenuItem("Help", nil)

	//Create menus: File, Help
	menuFile := fyne.NewMenu("File", menuItemNewProject, menuItemOpenProject, menuItemExportProject)
	menuHelp := fyne.NewMenu("Help", menuItemReport, menuItemSettings, menuItemHelp)

	//Set onto application window
	myWindow.SetMainMenu(fyne.NewMainMenu(menuFile, menuHelp))

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

	// Button to allow user to select a directory
	openProjectBtn := widget.NewButton("Open Project", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil || err != nil {
				return
			}
			state.SelectedDirectoryPath = uri.Path() // Store the selected directory path globally

			if _, err := os.Stat(state.SelectedDirectoryPath + "/BGM.ini"); err != nil {
				dialog.ShowError(errors.New("the selected folder is not a BGM Project folder"), myWindow)
				return
			}

			homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
			myWindow.SetContent(homeView)
		}, myWindow)
		menuItemExportProject.Disabled = false
	})

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

	newProjectBtn := widget.NewButton("New Project", func() {
		utils.PromptForProjectName(myWindow, func(newPath string) {
			state.SelectedDirectoryPath = newPath // Update the global variable with the new path
			//templatePath := "dialogue_temp.txt"
			//components.MakeNewFile(templatePath, state.SelectedDirectoryPath, myWindow)

			homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
			myWindow.SetContent(homeView) // Update the window content with the new tree
		})
		menuItemExportProject.Disabled = false
	})

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=
	// Initial content with the select directory button
	initialContent := container.NewVBox(
		layout.NewSpacer(),
		content,
		newProjectBtn,
		openProjectBtn,
		layout.NewSpacer(),
	)

	paddedContent := container.NewHBox(
		layout.NewSpacer(),
		initialContent,
		layout.NewSpacer(),
	)

	myWindow.SetContent(paddedContent)
	myWindow.Resize(fyne.NewSize(1280, 800))
	myWindow.ShowAndRun()
}
