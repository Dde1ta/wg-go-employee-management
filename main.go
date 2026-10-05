package main

import (
	"fmt"
	"os"

	"wg.dde1ta/pages"
)

func main() {
	fmt.Println("Welcome to the Employee Management System")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Login")
		fmt.Println("2. Employee Signup")
		fmt.Println("3. System Setup (Create Admin)")
		fmt.Println("4. Exit")
		fmt.Print("Select an option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			err := pages.LoginPage()
			if err != nil {
				fmt.Println("Returning to main menu...")
			}
		case "2":
			err := pages.EmployeeSignupPage()
			if err != nil {
				fmt.Println("Returning to main menu...")
			}
		case "3":
			err := pages.InitPage()
			if err != nil {
				fmt.Println("Returning to main menu...")
			}
		case "4":
			fmt.Println("Exiting system. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}