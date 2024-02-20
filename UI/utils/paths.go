package utils

import "path/filepath"

var parentDirectory = ""
var dialogueDirectory = ""
var scriptDirectory = parentDirectory + "/Script"
var installationDirectory = parentDirectory + "Installation"
var translationDirectory = parentDirectory + "/Translation"

func SetParentDirectory(directory string) {
	parentDirectory = directory
	dialogueDirectory = filepath.Join(parentDirectory, "Dialogue")
	scriptDirectory = parentDirectory + "/Script"
	installationDirectory = parentDirectory + "/Installation"
	translationDirectory = parentDirectory + "/Translation"
}

func GetParentDirectory() string {
	return parentDirectory
}

func GetDialogueDirectory() string {
	return dialogueDirectory
}

func GetScriptDirectory() string {
	return scriptDirectory
}

func GetInstallationDirectory() string {
	return installationDirectory
}

func GetTranslationDirectory() string {
	return translationDirectory
}
