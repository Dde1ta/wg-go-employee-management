package global

import (
	"errors"
	"fmt"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func TakeSecureInput(prompt string) (string, error) {
	var password string

	fmt.Print(prompt)

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nError reading password: %v\n", err)
		return "", nil
	}

	password = string(bytePassword)

	fmt.Println("Please Enter Again")
	fmt.Print(prompt)

	bytePassword, err = term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nError reading password: %v\n", err)
		return "", nil
	}

	if password != string(bytePassword){
		return "", errors.New("Passwords did not match")
	}
	
	return password, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
