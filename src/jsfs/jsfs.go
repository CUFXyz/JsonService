package jsfs

import (
	"fmt"
	"os"
	"strings"
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
	cssfiles, err := os.ReadDir("static/css")
	filename := strings.Split(hf.Name, ".")
	if len(cssfiles) == 0 || err != nil {
		fmt.Println("[FAILED | OPTIONAL]: STATIC FOLDER NOT FOUND OR CSS FILES NOT FOUND")
	} else {
		for _, file := range cssfiles {
			if file.Name() == fmt.Sprintf("%v.css", filename[0]) {
				hf.Css = file.Name()
				fmt.Println("[OK]: CSS FILE FOUNDED")
				break
			}
		}
	}
}

func checkJS(hf *Htmlfile) {
	jsfiles, err := os.ReadDir("static/js")
	filename := strings.Split(hf.Name, ".")
	if len(jsfiles) == 0 || err != nil {
		fmt.Println("[FAILED | OPTIONAL]: STATIC FOLDER NOT FOUND OR JS FILES NOT FOUND")
	} else {
		for _, file := range jsfiles {
			if file.Name() == fmt.Sprintf("%v.js", filename[0]) {
				hf.Js = file.Name()
				fmt.Println("[OK]: JS FILE FOUNDED")
				break
			}
		}
	}
}
