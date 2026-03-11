package main

import (
	_ "ShopAPI/docs"
	"ShopAPI/internal/database"
	"ShopAPI/internal/project"
	"ShopAPI/internal/user"

	httpSwagger "github.com/swaggo/http-swagger"

	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// @title Shop API
// @version 1.0
// @description API for freelance platform
// @host localhost:8000
// @BasePath /

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

	userStore := user.NewUser(db)
	userHandler := user.NewUserHandler(userStore)

	projectStore := project.NewProject(db)
	projectHandler := project.NewProjectHandle(projectStore)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	user.UserRouter(r, userHandler)
	project.ProjectRouter(r, projectHandler)

	log.Println("Server running on port:", port)

	log.Fatal(http.ListenAndServe(":"+port, r))

}
