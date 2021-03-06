package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/go-kit/kit/log/level"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"kwoc20-backend/controllers"
	"kwoc20-backend/models"
	logs "kwoc20-backend/utils/logs/pkg"
)

func initialMigration() {
	db, err := gorm.Open("sqlite3", "kwoc.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Project{})
}

func main() {

	initialMigration()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/oauth", controllers.UserOAuth).Methods("POST")
	router.HandleFunc("/mentor", controllers.MentorReg).Methods("POST")
	router.HandleFunc("/project", controllers.ProjectReg).Methods("POST")
	router.HandleFunc("/project/all", controllers.ProjectGet).Methods("GET")

	_ = level.Info(logs.Logger).Log("msg", fmt.Sprintf("Starting server on port "+port))

	error := http.ListenAndServe(":"+port,
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router))
	if error != nil {
		_ = level.Error(logs.Logger).Log("error", fmt.Sprintf("%v",error))
		os.Exit(1)
	}

	
}
