package jsfs

import (
	"fmt"
	"os"
)

type PageFile struct {
	name   string
	format string
	css    string
	js     string
}

func (m *PageFile) CheckPages() bool {
	var cssfile, jsfile bool
	if cssfiles, err := os.ReadDir(fmt.Sprintf("static/css")); err != nil {
		for i := 0; i < len(cssfiles); i++ {
			if file, err := os.Open(fmt.Sprintf("static/css/%v", m.name)); err != nil {
				m.css = file.Name()
				cssfile = true
			} else {
				fmt.Println("CSS FILE IS NOT FOUND: CHECK FAILED")
			}
		}
	}
}
