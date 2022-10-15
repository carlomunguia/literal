package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_readJSON(t *testing.T) {
	sampleJSON := map[string]interface{}{
		"foo": "bar",
	}

	body, _ := json.Marshal(sampleJSON)

	var decJSON struct {
		FOO string `json:"foo"`
	}

	req, err := http.NewRequest("POST", "/test", bytes.NewReader(body))
	if err != nil {
		t.Log(err)
	}

	reqRecorder := httptest.NewRecorder()
	defer req.Body.Close()

	err = testApp.readJSON(reqRecorder, req, &decJSON)
	if err != nil {
		t.Log(err)
	}
}

func Test_writeJSON(t *testing.T) {
	recRecorder := httptest.NewRecorder()
	payload := jsonResponse{
		Error:   false,
		Message: "test",
	}

	headers := make(http.Header)
	headers.Add("FOO", "BAR")
	err := testApp.writeJSON(recRecorder, http.StatusOK, payload, headers)
	if err != nil {
		t.Log(err)
	}

	testApp.environment = "production"
	err = testApp.writeJSON(recRecorder, http.StatusOK, payload, headers)
	if err != nil {
		t.Log(err)
	}
}
