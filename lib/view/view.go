package view

import (
	"os"
	"path/filepath"
	"runtime"
)

type View struct {
	html string
}

func getFilePath(path string) string {
	var _, b, _, _ = runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), path)
}

func (view *View) GetHtml() string {
	return view.html
}

func Get(filename string) *View {
	path := getFilePath("../../views/" + filename)
	html, _ := os.ReadFile(path)

	return &View{string(html)}
}
