package utils

func JSONtoWeiDuHelper(mod string, zipFolder string, files []string) {

	if mod == "dialogue" {
		JSONtoWeiDuDialogue(zipFolder, files)
	} else if mod == "script" {
		//JSONtoWeiDuScript(zipFolder, files)
	} else if mod == "installation" {
		//JSONtoWeiDuInstallation(zipFolder, files)
	} else {
		//JSONtoWeiDuTranslation(zipFolder, files)
	}
}

func JSONtoWeiDuDialogue(zipFolder string, files []string) {

}

func JSONtoWeiDuScript(zipFolder string, files []string) {

}

func JSONtoWeiDuInstallation(zipFolder string, files []string) {

}

func JSONtoWeiDuTranslation(zipFolder string, files []string) {

}
