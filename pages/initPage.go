package pages

import (
	"fmt"
	"wg.dde1ta/global"
)

func InitPage() error {
	var email string;

	fmt.Println("Enter the email: ")
	fmt.Scanln(&email)

	passwordPlainText, err := global.TakeSecureInput("Enter the password: ")

	if err != nil {
		fmt.Println("An error occured during input reading", err)
		return err
	}

	passwordHashed, err := global.HashPassword(passwordPlainText)

	if err != nil {
		fmt.Println("An error occured during hashing", err)
	}

	id, err := auth.CreateAdmin(email, passwordHashed)

	if err != nil {
		fmt.Println("An error occured during signup", err)
		return err
	}

	fmt.Println("Created New Admin with id:", id, "Please Proceed to login")
	return nil
}