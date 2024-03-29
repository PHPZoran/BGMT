package main

import (
	"UI/utils"
	"UI/views"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
	"os"
)

type AppState struct {
	SelectedDirectoryPath string
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func IsBGMDirectory(filename string) bool {

	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Baldur's Gate Mod Tool")
	state := &AppState{}
	content := widget.NewLabel("Select a directory...")

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

	//File -> Menu Options
	//Export Project Submenu Item
	menuItemExportProject := fyne.NewMenuItem("Export Project", func() {
		utils.WeiDuFileConversion(myWindow)

		//dialog.ShowInformation("Information", "WIP: Clicking this will display a confirmation to user.\n"+
		//	"Then it will trigger the backend to compile using the backend to be game ready.",
		//	myWindow)
	})

	//New Project Submenu Item
	menuItemNewProject := fyne.NewMenuItem("New Project", func() {
		utils.PromptForProjectName(myWindow, func(newPath string) {

			state.SelectedDirectoryPath = newPath // Update the global variable with the new path

			var tmp string
			tmpLength := len(utils.GetParentDirectory())
			if tmpLength == 0 {
				tmp = ""
			} else {
				tmp = utils.GetParentDirectory()
			}
			fmt.Println(utils.GetModFolder())
			utils.SetParentDirectory(newPath)
			fmt.Println(utils.GetParentDirectory())

			if !IsBGMDirectory(utils.GetParentDirectory() + "/BGM.ini") {
				homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
				myWindow.SetContent(homeView) // Update the window content with the new tree
				menuItemExportProject.Disabled = false
			} else {
				utils.SetParentDirectory(tmp)
				dialog.ShowInformation("Error", "The folder is already a BGM project folder.", myWindow)
			}
		})
	})
	//Open Project Submenu Item
	menuItemOpenProject := fyne.NewMenuItem("Open Project", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil || err != nil {
				return
			}
			state.SelectedDirectoryPath = uri.Path() // Store the selected directory path globally

			var tmp string
			tmpLength := len(utils.GetParentDirectory())
			if tmpLength == 0 {
				tmp = ""
			} else {
				tmp = utils.GetParentDirectory()
			}

			fmt.Println(utils.GetModFolder())
			utils.SetParentDirectory(uri.Path())
			fmt.Println(utils.GetParentDirectory())

			if IsBGMDirectory(utils.GetParentDirectory() + "/BGM.ini") {
				homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
				myWindow.SetContent(homeView)
			} else {
				utils.SetParentDirectory(tmp)
				dialog.ShowInformation("Error", "The folder is not a BGM project folder.", myWindow)
			}
		}, myWindow)
		menuItemExportProject.Disabled = false
	})

	//By default disable exporting projects
	menuItemExportProject.Disabled = true

	//Help -> Menu Options
	//Report Bug Submenu Item
	menuItemReport := fyne.NewMenuItem("Report Bug", func() {
		issueURL := "https://github.com/PHPZoran/BGMT/issues"

		labelText := widget.NewLabel("If you have encountered a bug or issue while running this application,\n" +
			"please submit an issue ticket by clicking the link below.\n" +
			"However, please check if your issue already exists prior to posting it.")
		labelText.Alignment = fyne.TextAlignCenter

		hyperlinkLabel := widget.NewHyperlink("Issue Ticket", &url.URL{})

		parsedURL, err := url.Parse(issueURL)
		if err != nil {
			// Handle error
			return
		}
		hyperlinkLabel.SetURL(parsedURL)

		hyperlinkLabel.OnTapped = func() {
			err := fyne.CurrentApp().OpenURL(parsedURL)
			if err != nil {
				log.Println(err)
			}
		}

		content := container.New(layout.NewVBoxLayout(),
			labelText,
			hyperlinkLabel,
		)

		dialog.ShowCustom("Report an Issue", "Close", content, myWindow)
	})
	//Settings Submenu Item
	menuItemSettings := fyne.NewMenuItem("Settings", func() {
		dialog.ShowInformation("Settings", "WIP: Clicking this will provide an application settings window", myWindow)
	})
	//Help Submenu Item
	menuItemHelp := fyne.NewMenuItem("Help", func() {
		dialog.ShowInformation("Help", "WIP: Clicking this will provide you guides for some insight into using this mod", myWindow)
	})

	//Create menus: File, Help (Quit added by default)
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

			var tmp string
			tmpLength := len(utils.GetParentDirectory())
			if tmpLength == 0 {
				tmp = ""
			} else {
				tmp = utils.GetParentDirectory()
			}

			utils.SetParentDirectory(uri.Path())
			fmt.Println(utils.GetParentDirectory())

			if IsBGMDirectory(utils.GetParentDirectory() + "/BGM.ini") {
				homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
				myWindow.SetContent(homeView)
				menuItemExportProject.Disabled = false
			} else {
				utils.SetParentDirectory(tmp)
				dialog.ShowInformation("Error", "The folder is not a BGM project folder.", myWindow)
			}

		}, myWindow)
		menuItemExportProject.Disabled = false
	})

	//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

	newProjectBtn := widget.NewButton("New Project", func() {
		utils.PromptForProjectName(myWindow, func(newPath string) {
			state.SelectedDirectoryPath = newPath // Update the global variable with the new path

			var tmp string
			tmpLength := len(utils.GetParentDirectory())
			if tmpLength == 0 {
				tmp = ""
			} else {
				tmp = utils.GetParentDirectory()
			}

			utils.SetParentDirectory(newPath)
			fmt.Println(utils.GetParentDirectory())

			if !IsBGMDirectory(utils.GetParentDirectory() + "/BGM.ini") {
				homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
				myWindow.SetContent(homeView) // Update the window content with the new tree
				menuItemExportProject.Disabled = false
			} else {
				utils.SetParentDirectory(tmp)
				dialog.ShowInformation("Error", "The folder is already a BGM project folder.", myWindow)
			}
		})
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
