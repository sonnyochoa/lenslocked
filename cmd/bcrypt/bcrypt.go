package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	switch os.Args[1] {
	case "hash":
		// hash the password
		hash(os.Args[2])
	case "compare":
		// compare the password with the hash
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing: %v", password)
		return
	}
	fmt.Printf("hashed password: %v\n", string(hashedBytes))
}

func compare(password, hash string) {
	fmt.Println(">>> TODO")
	fmt.Printf(">>> Compare the password: %q -> with the hash: %q\n", password, hash)
}
