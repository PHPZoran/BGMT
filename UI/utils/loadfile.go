package utils

import (
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"log"
)

func LoadFileContent(filename string, label *widget.Label) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read file:", err)
		label.SetText("Failed to load content")
	} else {
		label.SetText(string(content))
	}
}
