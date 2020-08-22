package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(mdp []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(mdp, bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return hash
}
