package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arturbaccarin/library-api-with-tests/handler"
)

func TestBookList(t *testing.T) {
	router := SetupRouter(handler.BookMock{})

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, httptest.NewRequest("GET", "/books", nil))

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected 200, but got %d", recorder.Code)
	}

	if recorder.Body.String() != "Hello" {
		t.Errorf("Expected Hello, but got %s", recorder.Body.String())
	}
}
