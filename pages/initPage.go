// initPage.go
package pages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wg.dde1ta/global"
)

func InitPage() error {
	exists, _ := auth.AdminExists()
	if exists {
		fmt.Println("An Admin user already exists. System setup is already complete.")
		return fmt.Errorf("admin already exists")
	}

	fmt.Println("\n--- System Setup ---")
	fmt.Println("No Admin user found. Create the first Admin.")

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
