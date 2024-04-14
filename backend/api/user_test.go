package api

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)


func TestCreateUser(t *testing.T) {
    body := strings.NewReader(`{"email": "test@example.com", "password": "password123"}`)
    req, err := http.NewRequest("POST", "/user/create", body)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateUser)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}
