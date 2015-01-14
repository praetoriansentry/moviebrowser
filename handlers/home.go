package handlers

import (
	"log"
	"net/http"
)

func Home(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Executing the home handler")
	var empty map[string]interface{}
	sendResponse(rw, rq, "home", empty)
}
