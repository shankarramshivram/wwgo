package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	ViewsDir  string = "/Users/sshivram/go-workspace/src/wwgo/views"
	LayoutDir string = ViewsDir + "/layouts"
	ViewsExt  string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	files = addViewsDirPrefix(files)
	files = addViewsExtSuffix(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func addViewsDirPrefix(files []string) []string {
	ret := make([]string, len(files))
	for i, f := range files {
		ret[i] = ViewsDir + "/" + f
	}
	return ret
}

func addViewsExtSuffix(files []string) []string {
	ret := make([]string, len(files))
	for i, f := range files {
		ret[i] = f + ViewsExt
	}
	return ret
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		panic(err)
	}
	return files
}
