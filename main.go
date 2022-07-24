package main

import (
	"context"
	"gows/handlers"
	"gows/middleware"
	"gows/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env failed")
	}
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseURL: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {

	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/product", handlers.InsertProductHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/product/{id}", handlers.GetProductByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/product/{id}", handlers.UpdateProductHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/product/{id}", handlers.DeleteProductHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/products", handlers.ListProductHandler(s)).Methods(http.MethodGet)

	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}
