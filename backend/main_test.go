package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {

	InitDB()
	defer db.Close()

	newUser := User{
		FirstName: "Toby",
		LastName:  "Maguire",
		Email:     "12345piderman@example.com",
		Password:  "password",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CreateUser returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedResponse := "New user was created"
	if rr.Body.String() != expectedResponse {
		t.Errorf("CreateUser returned unexpected response: got %v, want %v", rr.Body.String(), expectedResponse)
	}
}

func TestGetUsers(t *testing.T) {

	InitDB()
	defer db.Close()

	newUser := User{
		FirstName: "Toby",
		LastName:  "Maguire",
		Email:     "12345piderman@example.com",
		Password:  "password",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	req, err := http.NewRequest("GET", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetUsers returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func LogInUser(t *testing.T) { // NEW TEST

	InitDB()
	defer db.Close()

	newUser := User{
		FirstName: "Toby",
		LastName:  "Maguire",
		Email:     "12345piderman@example.com",
		Password:  "password",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	req, err := http.NewRequest("GET", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("LogIn returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

}

func TestUpdateUsers(t *testing.T) { // NEW TEST

	InitDB()
	defer db.Close()

	newUser := User{
		ID:        "59",
		FirstName: "Toby",
		LastName:  "Maguire",
		Email:     "12345piderman@example.com",
		Password:  "password",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", "/users/59", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetUsers returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedResponse := "User with ID =  was updated"
	if rr.Body.String() != expectedResponse {
		t.Errorf("CreateUser returned unexpected response: got %v, want %v", rr.Body.String(), expectedResponse)
	}
}

func TestGetUser(t *testing.T) { // NEW TEST

	InitDB()
	defer db.Close()

	newUser := User{
		ID:        "59",
		FirstName: "Toby",
		LastName:  "Maguire",
		Email:     "12345piderman@example.com",
		Password:  "password",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	req, err := http.NewRequest("GET", "/users/59", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetUsers returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestDeleteUser(t *testing.T) {

	InitDB()
	defer db.Close()

	newUser := User{
		ID:        "55",
		FirstName: "Toby",
		LastName:  "Maguire",
		Email:     "12345piderman@example.com",
		Password:  "password",
	}

	jsonUser, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Failed to marshal user to JSON: %v", err)
	}

	req, err := http.NewRequest("DELETE", "/users/55", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DeleteUser returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedResponse := "User with ID =  was deleted"
	if rr.Body.String() != expectedResponse {
		t.Errorf("CreateUser returned unexpected response: got %v, want %v", rr.Body.String(), expectedResponse)
	}
}
