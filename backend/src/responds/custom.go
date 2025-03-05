package responds

import (
	"log"
	"net/http"
)

func RespondDiceRoll(w http.ResponseWriter, code int, data interface{}) {
	if code > 499 {
		log.Printf("RespondWithError with 5XX error: (%d, %s)\n", code, data)
	}

	// diceRollResponse - ответ о броске кубиков
	type diceRollResponse struct {
		Error *string     `json:"error"`
		Data  interface{} `json:"data"`
	}

	ResponseWithJson(w, code, diceRollResponse{
		Error: nil,
		Data:  data,
	})
}
