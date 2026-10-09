// initPage.go
package pages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wg.dde1ta/entity"
	"wg.dde1ta/global"
)

func InitPage() error {

	exists, _ := admin.AdminExists()
	if exists {
		fmt.Println("Login as Admin")

		err := LoginSetUp()

		if err != nil {
			fmt.Println("Error Occured during login")
			return err
		}

		session, ok := global.GetGlobalSession()

		fmt.Println("Debug: ", session, "ok", ok)

		if !ok || session.UserRole != "setup" {
			return fmt.Errorf("login failed")
		}
	}else {
		global.NewSessionContext(&entity.Admin{Email: "InitSetupMode", Role: "setup", PasswordHashed: ""})
	}

	fmt.Println("\n--- Admin Setup ---")
	fmt.Println("Create the Admin.")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	passwordPlainText, err := global.TakeSignUpPassword("Enter the password: ")
	if err != nil {
		fmt.Println("An error occured during input reading", err)
		return err
	}

	passwordHashed, err := global.HashPassword(passwordPlainText)
	if err != nil {
		fmt.Println("An error occured during hashing", err)
	}

	err = auth.CreateAdmin(email, passwordHashed)
	if err != nil {
		fmt.Println("An error occured during signup", err)
		return err
	}

	fmt.Println("Created New Admin. Please Proceed to login")
	return nil
}
