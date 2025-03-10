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
	if cssfiles, err := os.ReadDir(fmt.Sprintf("static/css")); err != nil {
		for i := 0; i < len(cssfiles); i++ {

		}
	}
}
