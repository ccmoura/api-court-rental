package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		}

		server.DB.SingularTable(true)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("🚀 Listening to port " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(addr, server.Router))
}
