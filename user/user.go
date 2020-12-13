package user

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
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

func AddUser(username, password string) {
	log.Println("user 24")
	log.Println(username)
	log.Println(password)
	db := dbConn()
	insForm, err := db.Prepare("INSERT INTO user(username, password) VALUES(?,?)")
	if err != nil {
			panic(err.Error())
		}
		_, _ = insForm.Exec(username, password)
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
