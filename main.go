package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port_string := os.Getenv("PORT")
	if port_string == "" {
		log.Fatal("port not found")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port_string,
	}

	log.Printf("server is starting on %v", port_string)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
