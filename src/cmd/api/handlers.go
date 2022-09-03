package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type envelope map[string]interface{}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid /missing json"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
		return
	}

	app.infoLog.Println("Login request for", creds.Username)

	payload.Error = false
	payload.Message = "Login successful"

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
