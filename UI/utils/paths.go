package utils

import "path/filepath"

var parentDirectory string

func SetParentDirectory(directory string) { parentDirectory = directory }
func GetParentDirectory() string          { return parentDirectory }

func GetModFolder() string             { return filepath.Base(parentDirectory) }
func GetDialogueDirectory() string     { return parentDirectory + "/Dialogue" }
func GetScriptDirectory() string       { return parentDirectory + "/Scripts" }
func GetInstallationDirectory() string { return parentDirectory + "/Installation" }
func GetTranslationDirectory() string  { return parentDirectory + "/Translation/English" }

//func GetTemporaryDirectory() string { return parentDirectory + "/Templates" }
