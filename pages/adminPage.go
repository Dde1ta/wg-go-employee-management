// adminPage.go
package pages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"wg.dde1ta/global"
)

func AdminPage() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Admin Dashboard ---")
		fmt.Println("1. View All Employees")
		fmt.Println("2. Update Employee Details")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Logout")
		fmt.Print("Select an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			employees, err := admin.GetAllEmployees()
			if err != nil {
				fmt.Println("Error fetching employees:", err)
				continue
			}
			fmt.Println("\n--- Employee List ---")
			for _, emp := range employees {
				fmt.Println(emp)
			}
		case "2":
			fmt.Print("Enter Employee ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Enter field to update (department/position): ")
			field, _ := reader.ReadString('\n')
			field = strings.TrimSpace(field)

			fmt.Print("Enter new value: ")
			newData, _ := reader.ReadString('\n')
			newData = strings.TrimSpace(newData)

			err := admin.UpdateEmployeeDetails(id, newData, field)
			if err != nil {
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Employee updated successfully.")
			}
		case "3":
			fmt.Print("Enter Employee ID to delete: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			err := admin.DeleteEmployee(id)
			if err != nil {
				fmt.Println("Deletion failed:", err)
			} else {
				fmt.Println("Employee deleted successfully.")
			}
		case "4":
			global.LogOut()
			fmt.Println("Logged out successfully.")
			return nil
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
