package main

import (
	"encoding/json"
	"fmt"
	"github.com/lclarke98/go-project/user"
	"log"
	"net/http"
)

type UserDetails struct {
	Username string
	Password string
}

type DbResponse struct {
	Username string
	Password string
}


func register(req http.ResponseWriter, res *http.Request)  {
	var p UserDetails
	log.Println("here 24")
	_ = json.NewDecoder(res.Body).Decode(&p)
	log.Println(p.Username)
	log.Println(p.Password)

	log.Println("here 26")
	user.AddUser(string(p.Username), string(p.Password))

}

func main() {
	// add user route
	http.HandleFunc("/api/register", register)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
