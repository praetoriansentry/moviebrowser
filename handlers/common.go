package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	// I'm using two template libs. I don't need HTML escaping
	// when I'm pulling sub templates into the main layout. There
	// might be a better way to do this though..
	ttemplate "text/template"
)

var (
	// This is used to store all of the html templates in
	// memory. We don't need to go to disk every time we load a
	// page
	layouts    map[string]*template.Template
	baseLayout *ttemplate.Template
)

type (
	TemplateBody struct {
		Body         string
		TemplateName string
	}
)

// The visibility of this function is limited to the handlers
// package. It's used by the handlers in order to send a responds to
// the browser once all the data is prepared
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

// Very similar to the sendResponse function, but this is just
// specific to situations where we need to push Json objects
func marshalToJsonAndSend(rw http.ResponseWriter, data interface{}) {
	rw.Header().Add("Content-type", "application/json")
	json, err := json.Marshal(data)
	if err != nil {
		log.Printf("There was an issue marshaling to JSON: %q", err)
	}
	rw.Write(json)
}

// init functions run when the code is loaded. This function will
// load the base layout
func init() {
	log.Println("Loading templates")
	layouts = make(map[string]*template.Template)
	baseLayout = ttemplate.Must(ttemplate.ParseFiles("./templates/layout.html"))
}
