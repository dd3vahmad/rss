package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	router.Use(cors.Handler(cors.Options{ 
		AllowedOrigins: []string{"https://*", "http://*"}, 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

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