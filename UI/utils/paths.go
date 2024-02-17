package utils

var parentDirectory = ""
var dialogueDirectory = parentDirectory + "/Dialogue"
var scriptDirectory = parentDirectory + "/Script"
var installationDirectory = parentDirectory + "Installation"
var translationDirectory = parentDirectory + "/Translation"

func SetParentDirectory(directory string) {
	parentDirectory = directory
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
