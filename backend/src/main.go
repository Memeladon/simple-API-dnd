package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"simple-API-dnd/src/handlers"
	"simple-API-dnd/src/routers"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	// Фабрика роутеров
	factory := routers.NewRouterFactory()

	//customFactory := routers.NewRouterFactory(
	//	routers.WithAllowedOrigins([]string{"https://google.com"}),
	//	routers.WithAllowedMethods([]string{"GET", "POST"}),
	//)

	//Главный роутер
	mainRouter := factory.CreateRouter()

	//Версионные роутеры:
	apiDiceRouter := factory.CreateRouter()
	mainRouter.Mount("/api", apiDiceRouter)
	apiDiceRouter.Get("/health", handlers.HandlerLiveness)
	apiDiceRouter.Get("/err", handlers.HandlerError)
	apiDiceRouter.Get("/d/{n}", handlers.HandlerDiceRoll)

	// Запуск сервера
	srv := &http.Server{
		Handler: mainRouter,
		Addr:    ":" + portString,
	}

	log.Printf("Starting server on port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}
