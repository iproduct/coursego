package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

var rwlock sync.RWMutex
var database = make(map[int]User)
var sequence uint64

func SendError(w http.ResponseWriter, status int, err error, message string) {
	text := fmt.Sprintf("%s: %s", message, err)
	log.Println(text)
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"error": %s}`, text)
}

func users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		//buf := new(bytes.Buffer)
		//buf.ReadFrom(r.Body)
		//fmt.Printf("%s\n", buf)

		//body, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	SendError(w, http.StatusBadRequest,  err, "Error reading request body",)
		//	return
		//}
		user := User{}
		//if err := json.Unmarshal(buf.Bytes(), &user); err != nil {
		//if err := json.Unmarshal(body, &user); err != nil {
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			SendError(w, http.StatusBadRequest, err, "JSON unmarshaling failed")
			return
		}
		log.Printf("AFTER UNMARSHAL:%#v\n", user)
		newID := int(atomic.AddUint64(&sequence, 1))
		user.Id = newID
		user.Created = time.Now()
		user.Modified = user.Created
		rwlock.Lock()
		database[newID] = user
		rwlock.Unlock()
		w.Header().Add("Content Type", "application/json")
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(newID))
		w.WriteHeader(http.StatusCreated)

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "    ")
		err := encoder.Encode(&user)
		if err != nil {
			SendError(w, http.StatusInternalServerError, err, "JSON marshaling failed")
			return
		}
	case http.MethodGet:
		w.Header().Add("Content Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "    ")
		users := make([]User, 0, len(database))
		rwlock.RLock()
		for _, u := range database {
			users = append(users, u)
		}
		rwlock.RUnlock()

		err := encoder.Encode(users)
		//data, err := json.MarshalIndent(users, "", "    ")
		if err != nil {
			log.Printf("JSON marshaling failed: %s", err)
			SendError(w, http.StatusInternalServerError, err, "JSON marshaling failed")
			return
		}
		//w.Write(data)
	}
}

func main() {
	http.HandleFunc("/users", users)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
