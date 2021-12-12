package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursegopro/09-rest/rest_tdd/rest"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a rest.App

func TestMain(m *testing.M) {
	a = rest.App{}
	a.Initialize("root", "root", "go_rest_api")

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec("CREATE DATABASE IF NOT EXISTS `go_rest_api`"); err != nil {
		log.Fatal(err)
	}
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    age INT NOT NULL
)`

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/users/45", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test user", "email":"test@mydomain.com", "age":30}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)
	log.Println("Response:", string(response.Body.Bytes()))

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test user" {
		t.Errorf("Expected user name to be 'test user'. Got '%v'", m["name"])
	}
	if m["email"] != "test@mydomain.com" {
		t.Errorf("Expected email to be 'test@mydomain.com'. Got '%v'", m["name"])
	}

	if m["age"] != 30.0 {
		t.Errorf("Expected user age to be '30'. Got '%v'", m["age"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
	}
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", ("User " + strconv.Itoa(i+1)), ((i + 1) * 10))
		a.DB.Exec(statement)
	}
}

func TestGetUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

