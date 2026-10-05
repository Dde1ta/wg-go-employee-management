package main

import (
	"fmt"
	"wg.dde1ta/repo"
)


func main() {
	var employees repo.EmployeeRepo = *repo.NewEmployeeRepo("data/db.json")

	employeeOneId := "01a10a9c-c9fd-70c5-ad74-8ed06230416d"

	err := employees.DeleteEmployee(employeeOneId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(employees.GetById(employeeOneId))		
}