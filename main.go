package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not loaded:", err)
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set")
	}

	fmt.Println("PORT:", portString)

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Println("Server started on port", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
