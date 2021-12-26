package render

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
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
	case "jet":
		return render.JetPage(w, r, view, variables, data)
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

func (render *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	templateData := &TemplateData{}
	if data != nil {
		templateData = data.(*TemplateData)
	}

	t, err := render.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, templateData); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
