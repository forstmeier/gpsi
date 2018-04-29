package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_mainHandler(t *testing.T) {
	tests := []struct {
		request string
		status  int
	}{
		{"POST", http.StatusBadRequest},
	}

	for i := range tests {
		req, err := http.NewRequest(tests[i].request, "/unit-testing", nil)
		if err != nil {
			t.Errorf("error generating test request: %v", err)
		}
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(mainHandler)
		handler.ServeHTTP(rec, req)

		if status := rec.Code; status != tests[i].status {
			t.Errorf("handler returning wrong status, expected %v, received %v", tests[i].status, status)
		}
	}
}

func TestNew(t *testing.T) {
	output := New()
	if output == nil {
		t.Error("no server object returned from New()")
	}
}
