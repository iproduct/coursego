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

// error handling middleware
type AppError struct {
	Error   error
	Message string
	Code    int
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil { // err is *appError, not os.Error.
		SendError(w, err.Code, err.Error, "Application error: ") // or http.Error(w, err.Message, err.Code)
	}
}

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

func users(w http.ResponseWriter, r *http.Request) *AppError {
	log.Println("In request handler ...")
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()
		user := User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			return &AppError{Error: err, Message: "JSON unmarshaling failed", Code: http.StatusBadRequest}
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
			return &AppError{Error: err, Message: "JSON marshaling failed", Code: http.StatusInternalServerError}
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
			return &AppError{Error: err, Message: "JSON marshaling failed", Code: http.StatusInternalServerError}
		}
		w.Write(data)
	}
	return nil
}

func myMiddleware(handler AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *AppError {
		log.Println("Executing myMmiddleware before request...")
		// pass the call to the handler
		handler.ServeHTTP(w, r)
		log.Println("Executing myMmiddleware after request...")
		return nil
	}
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Main logic and response generation is implemented here
	log.Println("Executing myHandler...")
	w.Write([]byte("Response: OK"))
}

// Middleware: filter requests by application/json MIME type
func filterPOSTByContentType(handler AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *AppError {
		log.Println("Middleware 1: Filtering JSON MIME type ...")
		if r.Method == http.MethodPost && r.Header.Get("Content-type") != "application/json" {
			return &AppError{Error: nil, Message: "415 – Unsupported Media Type. Service accepts 'application/json' only.", Code: http.StatusUnsupportedMediaType}
		}
		handler.ServeHTTP(w, r)
		return nil
	}
}

// Middleware: add server timestamp as response header
func setServerTimeHeader(handler AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *AppError {
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
		w.Header().Add("Server-Time(UTC)", strconv.FormatInt(time.Now().Unix(), 10))
		w.WriteHeader(resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
		// Setting Server-Time header for all responses
		log.Println("Middleware 2: Setting server time custom header ...")
		return nil
	}
}

type Adapter func(handler AppHandler) AppHandler

func Apply(h AppHandler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return AppHandler(h)
}

func main() {
	http.Handle("/users", Apply(users, setServerTimeHeader, filterPOSTByContentType, myMiddleware))
	log.Fatal(http.ListenAndServe(":8088", nil))
}
