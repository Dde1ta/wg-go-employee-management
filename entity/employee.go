package entity

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"contact_number" validate:"required,len=10"`
	Role           string `json:"role"`
	PasswordHashed string `json:"password" validate:"required"`
	Department     string `json:"department" validate:"required"`
	Position       string `json:"position" validate:"required"`
}

func NewEmployee(name, email, password, phone, department, position string) (*Employee, error) {
	newEmployee := Employee{
		Name:           name,
		Email:          email,
		Phone:          phone,
		PasswordHashed: password,
		Department:     department,
		Position:       position,
		Role:           "employee",
	}

	err := myValidator.Struct(newEmployee)

	if err != nil {
		return nil, err
	}

	return &newEmployee, nil
}

func (emp *Employee) ToJson() (string, error) {
	jsonBytes, err := json.Marshal(emp)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (emp *Employee) GetProperty(property string) (Serializeable, error) {

	var value *MyString

	switch property {
	case "email":
		value.CopyString(emp.Email)
	case "role":
		value.CopyString(emp.Role)
	case "password":
		value.CopyString(emp.PasswordHashed)
	case "contact":
		value.CopyString(emp.Phone)
	case "department":
		value.CopyString(emp.Department)
	case "position":
		value.CopyString(emp.Position)
	case "name":
		value.CopyString(emp.Name)
	default:
		return nil, fmt.Errorf("The property %s is not valid for an employee", property)
	}
	return value, nil
}

func (emp *Employee) SetProperty(key string, value Serializeable) error {

	empCopy := *emp

	switch key {
	case "email":
		empCopy.Email = value.ToString()
	case "role":
		empCopy.Role = value.ToString()
	case "password":
		empCopy.PasswordHashed = value.ToString()
	case "contact":
		empCopy.Phone = value.ToString()
	case "department":
		empCopy.Department = value.ToString()
	case "position":
		empCopy.Position = value.ToString()
	case "name":
		empCopy.Name = value.ToString()
	default:
		return fmt.Errorf("The property %s is not valid for an employee", key)
	}

	err := empCopy.Validate()

	if err != nil {
		return err
	}

	emp = &empCopy

	return nil
}

func (emp *Employee) Validate() error {
	return myValidator.Struct(emp)
}

func (emp *Employee) String() string {
	return fmt.Sprintf("Employee Name: %s | Email: %s | Contact Number: %s | Department: %s | Position: %s", emp.Name,
		emp.Email,
		emp.Phone,
		emp.Department,
		emp.Position)
}
