package main

import (
	"fmt"
	"log/slog"
	"os"

	"wg.dde1ta/pages"
)

func main() {
	
	logFile, err := os.OpenFile("data/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	
	if err != nil {
		textLogger := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})
		logger := slog.New(textLogger)
		slog.SetDefault(logger)

		slog.Error("Log File cannot be opened logging to Stdout")
	}else{
		textLogger := slog.NewTextHandler(logFile, &slog.HandlerOptions{Level: slog.LevelDebug})
		logger := slog.New(textLogger)
		slog.SetDefault(logger)

		slog.Info("Log File Set")
	}

	slog.Info("Starting System")

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
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
