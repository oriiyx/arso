package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/oriiyx/arso/db"
	"github.com/oriiyx/arso/handler"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			log.Fatal(err)
		}
	})

	router.Get("/api/max-strength", handler.Make(handler.HandleMaxStrength))

	port := os.Getenv("HTTP_LISTEN_ADDR")

	slog.Info("Starting server on: ", "port", port)

	log.Fatal(http.ListenAndServe(port, router))

}

func initEverything() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := db.Init(); err != nil {
		return err
	}

	return nil
}
