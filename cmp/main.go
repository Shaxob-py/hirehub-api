package main

import (
	"ShopAPI/internal/database"
	"ShopAPI/internal/user"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	if port == "" {
		port = "8000"
	}

	db, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbTask := user.NewUser(db)
	r := chi.NewRouter()

	userHandler := user.NewUserHandler(dbTask) // if you have constructor
	user.UserRouter(r, userHandler)

}
