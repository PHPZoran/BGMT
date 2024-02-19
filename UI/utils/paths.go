package utils

var parentDirectory string

func SetParentDirectory(directory string) { parentDirectory = directory }

func GetParentDirectory() string {
	return parentDirectory
}
func GetDialogueDirectory() string { return parentDirectory + "/Dialogue" }
func GetScriptDirectory() string {
	return parentDirectory + "/Script"
}
func GetInstallationDirectory() string {
	return parentDirectory + "/Installation"
}
func GetTranslationDirectory() string {
	return parentDirectory + "/Translation"
}
