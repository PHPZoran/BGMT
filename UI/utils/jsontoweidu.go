package utils

func JSONtoWeiDuHelper(mod string, files []string) {
	if mod == "dialogue" {
		JSONtoWeiDuDialogue(files)
	} else if mod == "script" {
		JSONtoWeiDuScript(files)
	} else if mod == "installation" {
		JSONtoWeiDuInstallation(files)
	} else {
		JSONtoWeiDuTranslation(files)
	}
}

func JSONtoWeiDuDialogue(files []string) {
	
}

func JSONtoWeiDuScript(files []string) {

}

func JSONtoWeiDuInstallation(files []string) {

}

func JSONtoWeiDuTranslation(files []string) {

}
