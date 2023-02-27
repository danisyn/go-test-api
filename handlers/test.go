package handlers

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"api/database"
	"time"

	_ "github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
	Token string `json:"token"`
	Password string `json:"password"`
}

var status []string

func GetTest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	exists := database.GetUser(user.Name, user.Password, user.Token)
	if exists != false {
		fmt.Fprint(w,"True")
		status = append(status, "True")
	} else {
		fmt.Fprint(w,"False")
		status = append(status, "False")
	}
	
}

func Logs(w http.ResponseWriter, r *http.Request){
	body := " <html><head><title>The Tudors</title><meta http-equiv="+"refresh"+" content="+"1"+" /></head><h1>LOGS</h1>"
	fmt.Fprintln(w,body)
	for _ , v := range status {
		now := time.Now()
		time := fmt.Sprint(now)
		fmt.Fprintln(w, "<p>"+ time + "" +v+"</p>")
	}
	
}
