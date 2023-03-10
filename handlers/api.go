package handlers

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"api/database"

	_ "github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
	Token string `json:"token"`
	Password string `json:"password"`
}


func GetTest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	exists := database.GetUser(user.Name, user.Password, user.Token)
	if exists != false {
		fmt.Fprint(w,"True")
	} else {
		fmt.Fprint(w,"False")
	}
	
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	exists := database.GetUser(user.Name, user.Password, user.Token)
	if exists != false {
		_, response := database.GetConfig(user.Token, exists)
		fmt.Fprint(w, response)
	}
}
