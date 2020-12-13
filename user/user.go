package user

import (
	"database/sql"
	"encoding/json"
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

func addUser(username, password string) {
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
