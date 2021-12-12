package rest_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursego/10-modules-rest-jwtauth/model"
	"github.com/iproduct/coursego/10-modules-rest-jwtauth/rest"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a rest.App
var db *sql.DB

func init() {
	connectionString := fmt.Sprintf("%s:%s@/%s", "root", "root", "go_rest_api_test")
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	a = rest.App{}
	a.Init("root", "root", "go_rest_api_test")
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
	assertResponseCode(t, http.StatusOK, resp.Code)
	if body := resp.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s\n", body)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users/45", nil)
	response := executeRequest(req)
	assertResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestGetUserById(t *testing.T) {
	clearTable()
	testUser := model.User{1, "User 1", "user1@mydomain.com", "user1", 20, true}
	addUser(&testUser)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	log.Println("Response:", string(response.Body.Bytes()))

	assertResponseCode(t, http.StatusOK, response.Code)
	var actual model.User
	json.Unmarshal(response.Body.Bytes(), &actual)
	testUser.Password = ""
	assertEqual(t, testUser, actual)
}

func TestCreateUser(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test user","email":"test@gmail.com","password":"test123","age":30}`)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(payload))
	response := executeRequest(req)

	assertResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test user" {
		t.Errorf("Expected user name to be 'test user'. Got '%v'", m["name"])
	}

	if m["email"] != "test@gmail.com" {
		t.Errorf("Expected user email to be 'test@gmail.com'. Got '%v'", m["email"])
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

func TestUpdateUser(t *testing.T) {
	clearTable()
	testUser := model.User{1, "User 1", "user1@mydomain.com", "user1", 20, true}
	addUser(&testUser)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	var originalUser model.User
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	editedUser := originalUser
	editedUser.Name = "New User"
	editedUser.Password = "new_password"
	editedUser.Age = 42
	editedUser.Active = false

	payload, err := json.Marshal(&editedUser)
	if err != nil {
		t.Fatal(err)
	}
	req, _ = http.NewRequest("PUT", "/users/1", bytes.NewBuffer(payload))
	response = executeRequest(req)
	log.Println("Response:", string(response.Body.Bytes()))

	assertResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != float64(editedUser.ID) {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalUser.ID, m["id"])
	}

	if m["name"] != editedUser.Name {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalUser.Name, editedUser.Name, m["name"])
	}

	if m["age"] != float64(editedUser.Age) {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%[2]v'", originalUser.Age, editedUser.Age, m["age"])
	}

	if m["active"] != editedUser.Active {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%[2]v'", originalUser.Active, editedUser.Active, m["active"])
	}
}

func TestDeleteUser(t *testing.T) {
	clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	assertResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	response = executeRequest(req)

	assertResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	response = executeRequest(req)
	assertResponseCode(t, http.StatusNotFound, response.Code)
}

//Utility funcs
func assertResponseCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %[1]T: %[1]v. Got: %[2]v\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		n := strconv.Itoa(i + 1)
		addUser(&model.User{i, "User " + n, "user" + n + "@mydomain.com", "user" + n, (i + 2) * 10, true})
	}
}

func addUser(user *model.User) (*model.User, error) {
	statement := "INSERT INTO users(name, email, password, age, active) VALUES(?, ?, ?, ?, ?)"
	result, err := db.Exec(statement, user.Name, user.Email, user.Password, user.Age, user.Active)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	user.ID = int(id)
	//err = r.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ensureTableExists() {
	if _, err := db.Exec(databaseCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(tableDropSql); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(emailIndexDropQuery); err != nil {
		log.Println(err)
	}
	if _, err := db.Exec(emailIndexCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	db.Exec("DELETE FROM users")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const databaseCreationQuery = "CREATE DATABASE IF NOT EXISTS `go_rest_api_test` "

const tableDropSql = `DROP TABLE IF EXISTS users `
const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    active BOOLEAN                           
)`

const emailIndexDropQuery = `DROP INDEX uidx_email ON users `
const emailIndexCreationQuery = `CREATE UNIQUE INDEX uidx_email ON users (email)`
