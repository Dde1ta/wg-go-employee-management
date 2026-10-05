// login.go
package pages

import (
	"fmt"
	"wg.dde1ta/global"
)

func LoginPage() error {
	var email string

	fmt.Print("Enter the email: ")
	fmt.Scanln(&email)

	passwordPlainText, err := global.TakeSecureInput("Enter the password: ")

	if err != nil {
		fmt.Println("An error occured during input reading", err)
		return err
	}

	err = auth.Login(email, passwordPlainText)
	if err != nil {
		fmt.Println("Login failed:", err)
		return err
	}

	session, ok := global.GetGlobalSession()
	if !ok {
		return fmt.Errorf("failed to load session")
	}

	fmt.Printf("\nWelcome %s! Logged in as %s\n", session.UserEmail, session.UserRole)

	switch session.UserRole {
	case "admin":
		return AdminPage()
	case "employee":
		return EmployeePage()
	default:
		return fmt.Errorf("unknown user role")
	}
}