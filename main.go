package go_project

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


func login(req http.ResponseWriter, res *http.Request)  {
	var p UserDetails
	err := json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		user.AddUser(string("Test"), string("pw123"))
	}
}

func main() {
	// add user route
	http.HandleFunc("/api/login", login)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
