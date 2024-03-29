package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"time"
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

func SendError(w http.ResponseWriter, status int, err error, message string) {
	var text string
	if err != nil {
		text = fmt.Sprintf("%s: %s", message, err)
	} else {
		text = fmt.Sprintf("%s", message)
	}
	log.Println(text)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	// http.Error() could be used as well
	fmt.Fprintf(w, `{"error": "%s"}`, text)
}

func users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		user := User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			SendError(w, http.StatusBadRequest, err, "JSON unmarshaling failed")
			return
		}
		fmt.Printf("AFTER UNMARSHAL:%#v\n", user)
		newID := int(atomic.AddUint64(&sequence, 1))
		user.Id = newID
		database[newID] = user
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(newID))
		w.WriteHeader(http.StatusCreated)

		data, err := json.MarshalIndent(user, "", "    ")
		if err != nil {
			SendError(w, http.StatusBadRequest, err, "JSON marshaling failed")
			return
		}
		w.Write(data)
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
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

func myMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing myMmiddleware before request...")
		// pass the call to the handler
		handler.ServeHTTP(w, r)
		log.Println("Executing myMmiddleware after request...")
	})
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Main logic and response generation is implemented here
	log.Println("Executing myHandler...")
	w.Write([]byte("Response: OK"))
}

// Middleware: filter requests by application/json MIME type
func filterPOSTByContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 1: Filtering JSON MIME type ...")
		if r.Method == http.MethodPost && r.Header.Get("Content-type") != "application/json" {
			SendError(w, http.StatusUnsupportedMediaType, nil, "415 – Unsupported Media Type. Service accepts 'application/json' only.")
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Middleware: add server timestamp as response header
func setServerTimeHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := httptest.NewRecorder()
		// Providing wrapper instead of original response writer
		handler.ServeHTTP(recorder, r)
		resp := recorder.Result()
		// copy original headers
		for k, v := range resp.Header {
			for _, h := range v {
				w.Header().Add(k, h)
			}
		}
		// add new header
		w.Header().Add("Server-Time", strconv.FormatInt(time.Now().Unix(), 10))
		w.WriteHeader(resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		w.Write(body)
		// Setting Server-Time header for all responses
		log.Println("Middleware 2: Setting server time custom header ...")
	})
}

func main() {
	usersHandler := http.HandlerFunc(users)
	http.Handle("/users", myMiddleware(filterPOSTByContentType(setServerTimeHeader(usersHandler))))
	// HandlerFunc returns a HTTP Handler
	myHandler := http.HandlerFunc(myHandlerFunc)
	http.Handle("/my", myMiddleware(myHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
