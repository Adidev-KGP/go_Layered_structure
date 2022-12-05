package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func TestCreateEntry(t *testing.T) {

	var jsonStr = []byte(`{"name":"Adidev","age":40,"location":"Bangaore"}`)

	req, err := http.NewRequest("POST", "/api/newuser", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetEntryByID(t *testing.T) {

	req, err := http.NewRequest("GET", "api/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	//req.URL.RawQuery += q.Encode()
	fmt.Println(req.URL.RawQuery)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":1,"name":"gopher","location":"India","age":25}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllEntry(t *testing.T) {

	req, err := http.NewRequest("GET", "api/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	//req.URL.RawQuery += q.Encode()
	fmt.Println(req.URL.RawQuery)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
}

