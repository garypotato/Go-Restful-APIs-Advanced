package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request: %s path: %s error: %s", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("not found: %s path: %s", r.Method, r.URL.Path)

	writeJSONError(w, http.StatusNotFound, "the requested resource could not be found")

}
