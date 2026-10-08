// employeePage.go
package pages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wg.dde1ta/global"
)

func EmployeePage() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		session, ok := global.GetGlobalSession()
		if !ok {
			return fmt.Errorf("invalid session")
		}

		fmt.Println("\n--- Employee Dashboard ---")
		fmt.Println("1. View My Profile")
		fmt.Println("2. Update My Profile")
		fmt.Println("3. Logout")
		fmt.Print("Select an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			emp, err := employee.GetEmployeeByEmail(session.UserEmail)
			if err != nil {
				fmt.Println("Error loading profile:", err)
				continue
			}
			fmt.Println("\n--- My Profile ---")
			fmt.Println(emp)
		case "2":
			fmt.Println("\nSelect field to update:")
			fmt.Println("1. Name")
			fmt.Println("2. Email")
			fmt.Println("3. Contact Number")
			fmt.Println("4. Password")
			fmt.Print("Enter choice (1-4): ")

			fieldChoice, _ := reader.ReadString('\n')
			fieldChoice = strings.TrimSpace(fieldChoice)

			var field string
			switch fieldChoice {
			case "1":
				field = "name"
			case "2":
				field = "email"
			case "3":
				field = "contact"
			case "4":
				field = "password"
			default:
				fmt.Println("Invalid choice, please try again.")
				continue
			}

			var newData string
			if field == "password" {
				plainText, err := global.TakeSignUpPassword("Enter new password: ")
				if err != nil {
					fmt.Println("Error reading input:", err)
					continue
				}
				newData, _ = global.HashPassword(plainText)
			} else {
				fmt.Print("Enter new value: ")
				newDataRaw, _ := reader.ReadString('\n')
				newData = strings.TrimSpace(newDataRaw)
			}

			err := employee.UpdateEmployeeDetails(session.UserEmail, newData, field)
			if err != nil {
				if err.Error() == "Session Change error"{
					global.LogOut()
					fmt.Println("Logged out successfully.")
					return nil
				}
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Profile updated successfully.")
			}
		case "3":
			global.LogOut()
			fmt.Println("Logged out successfully.")
			return nil
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
