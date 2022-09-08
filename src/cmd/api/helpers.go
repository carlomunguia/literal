package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	var output []byte

	if app.environment == "development" {
		out, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
		output = out
	} else {
		out, err := json.Marshal(data)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
		output = out
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header().Set(k, v[0])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(output)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var customErr error

	switch {
	case strings.Contains(err.Error(), "SQLSTATE 23505"): // duplicate key value violates unique constraint
		customErr = errors.New("duplicate key value violates unique constraint")
		statusCode = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 22001"): // value too long for type character varying(255)")
		customErr = errors.New("value too long for type character varying(255)")
		statusCode = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 23503"): // insert or update on table \"users\" violates foreign key constraint \"users_role_id_fkey\"")
		customErr = errors.New("insert or update on table \"users\" violates foreign key constraint \"users_role_id_fkey\"")
		statusCode = http.StatusForbidden
	default:
		customErr = err
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = customErr.Error()

	app.writeJSON(w, statusCode, payload)
}
