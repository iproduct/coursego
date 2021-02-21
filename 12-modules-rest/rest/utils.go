package rest

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithValidationError(fields validator.ValidationErrorsTranslations, w http.ResponseWriter) {
	//Create a new map and fill it
	response := make(map[string]interface{})
	response["status"] = "error"
	response["message"] = "validation error"
	response["errors"] = fields
	respondWithJSON(w, http.StatusUnprocessableEntity, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		//An error occurred processing the json
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occured internally"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
