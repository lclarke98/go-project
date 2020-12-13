package go_project

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	user "user"
	"github.com/lclarke98/go-project/user"
)

type UserDetails struct {
	Username string
	Password string
}

type DbResponse struct {
	Username string
	Password string
}

// db connection
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// insert user details
func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var p UserDetails
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {
		insForm, err := db.Prepare("INSERT INTO user(username, password) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(p.Username, p.Password)
	}
	defer db.Close()
}

func login(w http.ResponseWriter, r *http.Request) {
	var p UserDetails
	var d DbResponse
	err := json.NewDecoder(r.Body).Decode(&p)
	db := dbConn()

	err = db.QueryRow("SELECT username,password FROM user WHERE username = ? AND password = ?", p.Username, p.Password).Scan(&d.Username, &d.Password)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		log.Println("User not found")
	} else {
		log.Println(d.Username)
		log.Println(d.Password)
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
