package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursego/moduleslab/rest"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a rest.App
var db *sql.DB

func init() {
	connectionString := fmt.Sprintf("%s:%s@/%s", "root", "root", "go_rest_api")
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	a = rest.App{}
	a.Init("root", "root", "go_rest_api")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)

}

//Tests
func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
	if body := resp.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s\n", body)
	}
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test user","email":"test@gmail.com","password":"test123","age":30, "active": true}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test user" {
		t.Errorf("Expected user name to be 'test user'. Got '%v'", m["name"])
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



//Utility funcs
func checkResponseCode(t *testing.T, expected int, actual int) {
	if(expected != actual) {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func ensureTableExists() {
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(emailIndexCreationQuery); err != nil {
		log.Println(err)
	}
}

func clearTable() {
	db.Exec("DELETE FROM users")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    age INT NOT NULL,
    active BOOL DEFAULT TRUE
)`
const emailIndexCreationQuery = `CREATE UNIQUE INDEX uidx_email ON users (email)`

