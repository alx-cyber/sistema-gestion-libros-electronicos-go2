package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBookHandler(t *testing.T) {
	router := SetupRouter() // asegúrate que exista esta función

	jsonData := []byte(`{
		"id": 1,
		"title": "Libro Test",
		"author": "Autor Test"
	}`)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Se esperaba status 201, se obtuvo %d", rr.Code)
	}
}
