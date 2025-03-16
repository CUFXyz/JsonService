package jsfs

import (
	"fmt"
	"os"
)

type Htmlfile struct {
	Name string // name of the file
	Css  string // name of css file
	Js   string // name of js file
}

func ReadHTML(htmlfile *Htmlfile) []byte {
	file, _ := os.ReadFile(fmt.Sprintf("static/html/%v", htmlfile.Name))
	return file
}

func CheckPage(hf *Htmlfile) {
	checkCSS(hf)
	checkJS(hf)
}

func checkCSS(hf *Htmlfile) {
	cssfiles, _ := os.ReadDir("static/css")
	if len(cssfiles) == 0 {
		fmt.Println("[FAILED | OPTIONAL]: STATIC FOLDER NOT FOUND OR CSS FILES NOT FOUND")
	} else {
		for _, file := range cssfiles {
			if file.Name() == fmt.Sprintf("%v.css", hf.Name) {
				hf.Css = file.Name()
				fmt.Println("[OK]: CSS FILE FOUNDED")
				break
			}
		}
	}
}

func checkJS(hf *Htmlfile) {
	jsfiles, _ := os.ReadDir("static/js")
	if len(jsfiles) == 0 {
		fmt.Println("[FAILED | OPTIONAL]: STATIC FOLDER NOT FOUND OR JS FILES NOT FOUND")
	} else {
		for _, file := range jsfiles {
			if file.Name() == fmt.Sprintf("%v.js", hf.Name) {
				hf.Js = file.Name()
				fmt.Println("[OK]: JS FILE FOUNDED")
				break
			}
		}
	}
}
