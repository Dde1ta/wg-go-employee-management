package entity

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Employee struct {
	Id    string            `json:"id" validate:"required"`
	Name  string          	`json:"name" validate:"required"`
	Email string 		  	`json:"email" validate:"required,email"`
	Phone string 		  	`json:"phone_number" validate:"required,len=10"`
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
