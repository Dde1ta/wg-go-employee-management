// employeePage.go
package pages

import (
	"fmt"
	"wg.dde1ta/global"
)

func EmployeePage() error {
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

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			emp, err := employee.GetEmployeeById(session.UserId)
			if err != nil {
				fmt.Println("Error loading profile:", err)
				continue
			}
			fmt.Println("\n--- My Profile ---")
			fmt.Printf("ID: %s\nName: %s\nEmail: %s\nPhone: %s\nDepartment: %s\nPosition: %s\n", 
				emp.Id, emp.Name, emp.Email, emp.Phone, emp.Department, emp.Position)
		case "2":
			var field, newData string
			fmt.Print("Enter field to update (name/email/contact_number/password): ")
			fmt.Scanln(&field)

			if field == "password" {
				plainText, err := global.TakeSecureInput("Enter new password: ")
				if err != nil {
					fmt.Println("Error reading input:", err)
					continue
				}
				newData, _ = global.HashPassword(plainText)
			} else {
				fmt.Print("Enter new value: ")
				fmt.Scanln(&newData)
			}

			err := employee.UpdateEmployeeDetails(session.UserId, newData, field)
			if err != nil {
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