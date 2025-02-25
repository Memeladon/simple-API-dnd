package handlers

import (
	"math/rand"
	"net/http"
	"os"
	"simple-API-dnd/src/responds"
	"strconv"
	"strings"
)

type CheatType string

const (
	cheatLower      CheatType = "lower"
	cheatUpper      CheatType = "upper"
	quantitiesLimit           = 100
)

func HandlerDiceRoll(w http.ResponseWriter, r *http.Request) {

	correctToken := os.Getenv("TOKEN")
	cheats := r.Header.Get("X-CHEAT")
	authorization := r.Header.Get("Authorization")

	minLimit := 1
	quantity := 1

	if cheats == string(cheatLower) || cheats == string(cheatUpper) {
		if authorization == "" {
			responds.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		if strings.HasPrefix(authorization, "Bearer") {
			token := authorization[7:]
			if token != correctToken {
				responds.RespondWithError(w, http.StatusForbidden, "Unauthorized")
				return
			}
		} else {
			responds.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
	}

	// Получение количества сторон кубика
	diceFaces, err := strconv.Atoi(r.URL.Path[len("/api/d/"):])
	if err != nil {
		responds.RespondWithError(w, http.StatusBadRequest, "Invalid Dice Sides")
		return
	}

	if diceFaces < 4 {
		responds.RespondWithError(w, http.StatusBadRequest, "Dice Sides must be at least 4")
		return
	}

	// Проверка читов
	if cheats == string(cheatLower) {
		diceFaces = diceFaces / 2
	} else if cheats == string(cheatUpper) {
		minLimit = diceFaces / 2
	}

	// Получение количества кубиков
	quantityStr := r.URL.Query().Get("quantity")
	if quantityStr == "" {
		quantity = 1
	} else {
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			responds.RespondWithError(w, http.StatusBadRequest, "Invalid quantity")
			return
		}
	}
	if quantity <= 0 {
		responds.RespondWithError(w, http.StatusBadRequest, "Quantity must be positive")
		return
	}
	if quantity > quantitiesLimit {
		responds.RespondWithError(w, http.StatusBadRequest, "Quantity must be positive")
		return
	}

	// Обработка броска кубиков
	values := make([]int, quantity)
	for i := range values {
		values[i] = rand.Intn(diceFaces) + minLimit
	}

	// Данные о броске кубиков
	respond := struct {
		Values []int `json:"values"`
	}{Values: values}

	responds.RespondDiceRoll(w, http.StatusOK, respond)
}
