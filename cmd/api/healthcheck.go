package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	js := fmt.Sprintf(`{"status": "available", "environment": "%s", "version": "%s"}`, app.config.env, version)
	w.Header().Add("Content-Type", "application/json")

	w.Write([]byte(js))
}
