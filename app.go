package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "user/apiProject/config"
	. "user/apiProject/dao"
)

var config = Config{}
var dao = UserDAO{}

// Get list of users
func AllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Fuckers!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/users", AllUsersEndPoint).Methods("GET")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
