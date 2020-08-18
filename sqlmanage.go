package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnexionBDD(chemin string) *sql.DB {
	bdd, err := sql.Open("sqlite3", chemin)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return bdd

}

func Recuperationmotdepasse(bdd *sql.DB, username string) string {
	var requete string = "SELECT password FROM user where username = '" + username + "'"
	rows, err := bdd.Query(requete)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var password string
	for rows.Next() {
		rows.Scan(&password)
	}

	return password
}
