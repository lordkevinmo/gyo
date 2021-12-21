package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          string
}

func (render *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(render.Renderer) {
	case "go":
		return render.GoPage(w, r, view, data)
	}
	return nil
}

func (render *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%spage.tmpl", render.RootPath, view))
	if err != nil {
		return err
	}

	templateData := &TemplateData{}
	if data != nil {
		templateData = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &templateData)
	if err != nil {
		return err
	}

	return nil
}
