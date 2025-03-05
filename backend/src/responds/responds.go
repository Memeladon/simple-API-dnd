package responds

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("RespondWithError with 5XX error: (%d, %s)\n", code, message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	ResponseWithJson(w, code, errResponse{Error: message})
}

func ResponseWithJson(w http.ResponseWriter, code int, data interface{}) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dataJson)
}
