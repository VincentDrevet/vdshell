package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnexionBDD(chemin string) *sql.DB {
	bdd, err := sql.Open("sqlite3", chemin)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return bdd

}

func Recuperationmotdepasse(bdd *sql.DB, username string) string {
	var requete string = "SELECT password FROM user where username = '" + username + "'"
	rows, err := bdd.Query(requete)
	if err != nil {
		fmt.Printf("%v", err)

	}
	var password string
	for rows.Next() {
		rows.Scan(&password)
	}

	return password
}

func ChangementMotdePasse(bdd *sql.DB, hash []byte) {
	var requete string = "UPDATE user SET password = '" + string(hash) + "' WHERE username = 'admin'"

	stm, err := bdd.Prepare(requete)
	if err != nil {
		fmt.Printf("%v", err)
	}
	stm.Exec()

}
