package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Active   bool
}

var database = make(map[int]User)
var sequence uint64

func users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		//buf := new(bytes.Buffer)
		//buf.ReadFrom(r.Body)
		//fmt.Printf("%s\n", buf)
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user := User{}
		if err := json.Unmarshal(body, &user); err != nil {
		//if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Printf("JSON unmarshaling failed: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("AFTER UNMARSHAL:%#v\n", user)
		newID := int(atomic.AddUint64(&sequence, 1))
		user.Id = newID
		database[newID] = user
		w.Header().Add("Content Type", "application/json")
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(newID))
		w.WriteHeader(http.StatusCreated)

		data, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		w.Write(data)
	case http.MethodGet:
		w.Header().Add("Content Type", "application/json")
		users := make([]User, len(database))
		i := 0
		for _, u := range database {
			users[i] = u
			i++
		}
		data, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			log.Printf("JSON marshaling failed: %s", err)
		}
		w.Write(data)
	}
}

func main() {
	http.HandleFunc("/users", users)
	//http.HandleFunc("/headers", headers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
