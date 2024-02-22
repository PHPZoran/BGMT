package views

import (
	"UI/components"
	"UI/utils"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"path/filepath"
)

func MakeInstallationView(directoryPath string, window fyne.Window) fyne.CanvasObject {
	//Setting default variables
	newDirectoryPath := filepath.Join(directoryPath, "Installation")
	workingFilePath := filepath.Join(newDirectoryPath, "working.tmp")
	//defaultInstallationFilePath := filepath.Join(newDirectoryPath, "installation_example.txt")
	//skeletonInstallationFilePath := filepath.Join(newDirectoryPath, "installation_skeleton.txt")

	//Set Toolbar
	speakerID := ""
	modType := "Installation"
	extension := ".tp2"
	toolbar := CreateToolbar(directoryPath, window, speakerID, modType, extension, newDirectoryPath)

	//Setting the display box
	contentLabel := widget.NewLabel("Preview")
	contentLabel.Wrapping = fyne.TextWrapWord

	// Load and display the default file content

	//Buttons for Initial Dialogue options
	btnToNextInstallationPage := widget.NewButton("Next", func() {
		SetInstallationHeader(window)
		//NavigateTo(window, directoryPath, installation)
	})

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

	// Hide Next button until New or Load is clicked.
	btnToNextInstallationPage.Hide()

	btnForNewInstallation := widget.NewButton("New", func() {
		SetInstallationHeader(window)
		//components.MakeNewFile(templateInstallationFilePath, newDirectoryPath, window)
		utils.UpdateFileContent(workingFilePath)
		tree.Refresh() //TODO
		btnToNextInstallationPage.Show()
	})

	btnForLoadModFile := components.CreateLoadModButton(window, ".tp2", newDirectoryPath, func() {
		utils.UpdateFileContent(workingFilePath)
		tree.Refresh() //TODO
		btnToNextInstallationPage.Show()
	})

	paddedButtonBar2 := container.NewHBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
		btnToNextInstallationPage,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	padButtonBar2 := container.NewVBox(
		layout.NewSpacer(),
		paddedButtonBar2,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	// VBox containing your buttons and file preview
	btnOptions := container.NewVBox(
		layout.NewSpacer(),
		btnForNewInstallation,
		btnForLoadModFile,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	filePreview := container.NewVBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	grid := container.NewAdaptiveGrid(4,
		layout.NewSpacer(),
		btnOptions,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	vBox := container.NewVBox(
		layout.NewSpacer(),
		grid,
		padButtonBar2,
	)

	vBox2 := container.NewVBox(
		toolbar,
		widget.NewLabel(""),
		vBox,
	)

	hBox := container.NewHBox(vBox2, filePreview)

	split := container.NewHSplit(tree, hBox)
	split.Offset = .15

	return split

}
func SetInstallationHeader(window fyne.Window) {
	authorEntry := widget.NewEntry()
	languageEntry := widget.NewEntry()
	newDirectoryPath := utils.GetInstallationDirectory()
	templateInstallationFilePath := filepath.Join(newDirectoryPath, "installation_temp.txt")

	content := widget.NewForm(
		widget.NewFormItem("Author:", authorEntry),
		widget.NewFormItem("Language:", languageEntry),
	)

	dialog.ShowCustomConfirm("New Installation", "Confirm", "Cancel", content, func(confirmed bool) {
		if confirmed {
			author := authorEntry.Text
			language := languageEntry.Text
			fmt.Println("Author:", author)
			fmt.Println("Language:", language)
			components.MakeNewFile(templateInstallationFilePath, newDirectoryPath, window)
			//TODO: Insert author and language to new file and Save File
		}
	}, window)
}
