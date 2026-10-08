// login.go
package pages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wg.dde1ta/global"
)

func LoginPage() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	passwordPlainText, err := global.TakeLoginPassword("Enter the password: ")
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


func LoginSetUp() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	passwordPlainText, err := global.TakeLoginPassword("Enter the password: ")
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
	
	global.SetSetupRole()

	fmt.Printf("\nWelcome %s! Logged in as %s\n", session.UserEmail, session.UserRole)


	return nil
	
}