package main

import (
	"UI/utils"
	"UI/views"
	"archive/zip"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io"
	"os"
)

type AppState struct {
	SelectedDirectoryPath string
}

//-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=

// ZipFiles compresses one or more files into a zip archive.
// Credit: https://golangcode.com/create-zip-files-in-go/
func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err := addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

// addFileToZip adds a file to the zip archive
func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	// Create a header for the file
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Set the name of the file inside the zip file
	header.Name = filename

	// Add metadata to the file header if needed
	header.Method = zip.Deflate

	// Create a writer for the file in the zip archive
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Write the file data to the zip archive
	_, err = io.Copy(writer, fileToZip)
	return err
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
			utils.SetParentDirectory(newPath)
			fmt.Println(utils.GetParentDirectory())

			homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
			myWindow.SetContent(homeView) // Update the window content with the new tree
		})
		menuItemExportProject.Disabled = false
	})
	//Open Project Submenu Item
	menuItemOpenProject := fyne.NewMenuItem("Open Project", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri == nil || err != nil {
				return
			}
			state.SelectedDirectoryPath = uri.Path() // Store the selected directory path globally
			utils.SetParentDirectory(uri.Path())
			fmt.Println(utils.GetParentDirectory())

			if _, err := os.Stat(state.SelectedDirectoryPath + "/BGM.ini"); err != nil {
				dialog.ShowError(errors.New("the selected folder is not a BGM Project folder"), myWindow)
				return
			}

			homeView := views.MakeHomeView(state.SelectedDirectoryPath, myWindow)
			myWindow.SetContent(homeView)
		}, myWindow)
		menuItemExportProject.Disabled = false
	})

	//By default disable exporting projects
	menuItemExportProject.Disabled = true

	//Help -> Menu Options
	//Report Bug Submenu Item
	menuItemReport := fyne.NewMenuItem("Report Bug", func() {
		dialog.ShowInformation("Report an Issue",
			"If you have encountered a bug or issue while running this application,\n"+
				"please submit an issue ticket here: \n\nhttps://github.com/PHPZoran/BGMT/issues.\n\n"+
				"However, please check if your issue already exists prior to posting it.", myWindow)
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
			utils.SetParentDirectory(uri.Path())
			fmt.Println(utils.GetParentDirectory())

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
			utils.SetParentDirectory(newPath)
			fmt.Println(utils.GetParentDirectory())

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
