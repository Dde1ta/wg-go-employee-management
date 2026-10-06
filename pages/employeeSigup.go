// employeeSigup.go
package pages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wg.dde1ta/global"
)

func EmployeeSignupPage() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n--- Employee Signup ---")

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter Phone (10 digits): ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Enter Department: ")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	fmt.Print("Enter Position: ")
	position, _ := reader.ReadString('\n')
	position = strings.TrimSpace(position)

	passwordPlainText, err := global.TakeSignUpPassword("Enter a password: ")
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
