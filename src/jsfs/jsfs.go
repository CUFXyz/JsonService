package jsfs

import (
	"fmt"
	"os"
)

type Htmlfile struct {
	Name   string // name of the file
	Format string // format (need to be .html)
	Css    string // path to css file
	Js     string // path to js file
}

func CheckPage(hf *Htmlfile) {
	checkCSS(hf)
	checkJS(hf)
}

func checkCSS(hf *Htmlfile) {
	cssfiles, _ := os.ReadDir("static/css")
	if len(cssfiles) == 0 {
		fmt.Println("[FAILED | OPTIONAL]: STATIC FOLDER NOT FOUND OR CSS FILE NOT FOUND")
	} else {
		for _, file := range cssfiles {
			if file.Name() == fmt.Sprintf("%v.css", hf.Name) {
				fmt.Println("[OK]: CSS FILE FOUNDED")
				break
			}
		}
	}
}

func checkJS(hf *Htmlfile) {
	jsfiles, _ := os.ReadDir("static/js")
	if len(jsfiles) == 0 {
		fmt.Println("[FAILED | OPTIONAL]: STATIC FOLDER NOT FOUND OR JS FILE NOT FOUND")
	} else {
		for _, file := range jsfiles {
			if file.Name() == fmt.Sprintf("%v.js", hf.Name) {
				fmt.Println("[OK]: CSS FILE FOUNDED")
				break
			}
		}
	}
}
