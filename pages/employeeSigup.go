// employeeSigup.go
package pages

import (
	"fmt"
	"wg.dde1ta/global"
)

func EmployeeSignupPage() error {
	fmt.Println("\n--- Employee Signup ---")
	
	var name, email, phone, department, position string

	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter Email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter Phone (10 digits): ")
	fmt.Scanln(&phone)
	fmt.Print("Enter Department: ")
	fmt.Scanln(&department)
	fmt.Print("Enter Position: ")
	fmt.Scanln(&position)

	passwordPlainText, err := global.TakeSecureInput("Enter a password: ")
	if err != nil {
		fmt.Println("An error occured during input reading", err)
		return err
	}

	passwordHashed, err := global.HashPassword(passwordPlainText)
	if err != nil {
		fmt.Println("An error occured during hashing", err)
		return err
	}

	id, err := auth.SignUpEmployee(name, email, passwordHashed, phone, department, position)
	if err != nil {
		fmt.Println("Signup failed:", err)
		return err
	}

	fmt.Printf("Employee successfully created with ID: %s. Please login.\n", id)
	return nil
}