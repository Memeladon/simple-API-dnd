package handlers

import (
	"net/http"
	"simple-API-dnd/src/responds"
)

func HandlerLiveness(w http.ResponseWriter, r *http.Request) {
	responds.ResponseWithJson(w, http.StatusOK, struct{}{})
}

func HandlerError(w http.ResponseWriter, r *http.Request) {
	responds.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
