package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello world")

	godotenv.Load(".env")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("PORT cannot be found in environment")
	}

	fmt.Printf("PORT: %s", PORT)

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr: ":" + PORT,
	}

	log.Printf("Server running on port: %s", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}