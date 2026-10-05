package entity

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Employee struct {
	Id    string            `json:"id" validate:"required"`
	Name  string          	`json:"name" validate:"required"`
	Email string 		  	`json:"email" validate:"required,email"`
	Phone string 		  	`json:"contact_number" validate:"required,len=10"`
	Role  string			`json:"role"`		
	PasswordHashed string 	`json:"password" validate:"required"`
	Department     string 	`json:"department" validate:"required"`
	Position       string 	`json:"position" validate:"required"` 		
	 		
}

func NewEmployee(name, email, password, phone, department, position string) (Employee, error) {
	var newId, err = uuid.NewV7()

	if err != nil {
		return Employee{}, nil
	}

	newEmployee := Employee{
		Id: newId.String(),
		Name: name,
		Email: email,
		Phone: phone,
		PasswordHashed: password,
		Department: department,
		Position: position,
		Role: "employee",
	}

	err = myValidator.Struct(newEmployee)

	if err != nil {
		return Employee{}, err
	}

	return newEmployee, nil
}

func (emp Employee) ToJson() (string, error){
	jsonBytes, err := json.Marshal(emp)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (emp Employee) GetId() (string) {
	return emp.Id
}

func (emp Employee) GetRole() (string) {
	return emp.Role
}

func (emp Employee) GetEmail() (string) {
	return emp.Email
}

func (emp Employee) Validate() (error) {
	return myValidator.Struct(emp)
}

func (emp Employee) GetPassword() (string) {
	return emp.PasswordHashed
}
