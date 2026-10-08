package global

import (
	"fmt"
	"syscall"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

func TakeLoginPassword(prompt string) (string, error) {
	var password string

	fmt.Print(prompt)

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nError reading password: %v\n", err)
		return "", nil
	}

	password = string(bytePassword)

	return password, nil
}

// Helper function to check password strength
func isStrongPassword(s string) bool {
	if len(s) <= 8 {
		return false
	}
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, c := range s {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsNumber(c) || unicode.IsDigit(c):
			hasNumber = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}

func TakeSignUpPassword(prompt string) (string, error) {
	for {
		fmt.Print(prompt)

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("\nError reading password: %v\n", err)
			return "", err // Return the actual error
		}
		fmt.Println() // Add newline since ReadPassword suppresses it

		password := string(bytePassword)

		if !isStrongPassword(password) {
			fmt.Println("Password is too weak. It must be > 8 characters and include an uppercase letter, a lowercase letter, a number, and a special character.")
			fmt.Println("Let's try again.")
			continue
		}

		fmt.Print("Please Enter Again: ")

		bytePasswordConfirm, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("\nError reading password: %v\n", err)
			return "", err
		}
		fmt.Println()

		if password != string(bytePasswordConfirm) {
			fmt.Println("Passwords did not match. Let's try again.")
			continue
		}

		return password, nil
	}
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
