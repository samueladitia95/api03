package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	"encoding/json"
	"os"
	"fmt"
)

var db *sql.DB
var err error

// DBConfig type struct
type DBConfig struct {
	User string `json:"user"`
	DBEngine string `json:"dbengine"`
	Host string `json:"host"`
	Port int `json:"port"`
	DBName string `json:"dbname"`
}

func mainRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", Delivery).Methods("POST")

	log.Fatal(http.ListenAndServe(":8075", router))
}

func main() {
	dbConfigJSON, err := os.Open("./app/DBconfig.json")
	if err != nil {
		panic(err.Error())
	}

	var dbConfig DBConfig
	err = json.NewDecoder(dbConfigJSON).Decode(&dbConfig)
	if err != nil {
		panic(err.Error())
	}

	dbInfo := fmt.Sprintf("host:%s port:%d dbname:%s user:%s sslmode:disable", dbConfig.Host, dbConfig.Port, dbConfig.DBName, dbConfig.User)

	db, err = sql.Open(dbConfig.DBEngine, dbInfo)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	mainRouter()
}