package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	ttemplate "text/template"
)

var (
	layouts    map[string]*template.Template
	baseLayout *ttemplate.Template
)

type (
	TemplateBody struct {
		Body         string
		TemplateName string
	}
)

func sendResponse(rw http.ResponseWriter, rq *http.Request, templateName string, rawData map[string]interface{}) {
	rw.Header().Add("Content-type", "text/html")
	var err error
	if layouts[templateName] == nil {
		layouts[templateName], err = template.ParseFiles("./templates/" + templateName + ".html")
	}
	if err != nil {
		log.Printf("There was an issue loading the template named %s\n", templateName)
		http.NotFound(rw, rq)
		return
	}
	buf := bytes.NewBufferString("")
	layouts[templateName].Execute(buf, rawData)
	body := new(TemplateBody)
	body.Body = buf.String()
	body.TemplateName = templateName
	baseLayout.Execute(rw, body)
}

func marshalToJsonAndSend(rw http.ResponseWriter, data interface{}) {
	rw.Header().Add("Content-type", "application/json")
	json, err := json.Marshal(data)
	if err != nil {
		log.Printf("There was an issue marshaling to JSON: %q", err)
	}
	rw.Write(json)
}

func init() {
	log.Println("Loading templates")
	layouts = make(map[string]*template.Template)
	baseLayout = ttemplate.Must(ttemplate.ParseFiles("./templates/layout.html"))
}
