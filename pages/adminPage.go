// adminPage.go
package pages

import (
	"fmt"
	"wg.dde1ta/global"
)

func AdminPage() error {
	for {
		fmt.Println("\n--- Admin Dashboard ---")
		fmt.Println("1. View All Employees")
		fmt.Println("2. Update Employee Details")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Logout")
		fmt.Print("Select an option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			employees, err := admin.GetAllEmployees()
			if err != nil {
				fmt.Println("Error fetching employees:", err)
				continue
			}
			fmt.Println("\n--- Employee List ---")
			for _, emp := range employees {
				fmt.Printf("ID: %s | Name: %s | Email: %s | Dept: %s | Pos: %s\n", emp.Id, emp.Name, emp.Email, emp.Department, emp.Position)
			}
		case "2":
			var id, field, newData string
			fmt.Print("Enter Employee ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter field to update (department/position): ")
			fmt.Scanln(&field)
			fmt.Print("Enter new value: ")
			fmt.Scanln(&newData)

			err := admin.UpdateEmployeeDetails(id, newData, field)
			if err != nil {
				fmt.Println("Update failed:", err)
			} else {
				fmt.Println("Employee updated successfully.")
			}
		case "3":
			var id string
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scanln(&id)
			
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