package entity

import (
	"encoding/json"
	"fmt"
)

type Admin struct {
	Email          string `json:"email" validate:"required,email"`
	Role           string `json:"role"`
	PasswordHashed string `json:"password" validate:"required"`
}

func NewAdmin(email, hashedPassword string) (*Admin, error) {
	newAdmin := Admin{
		Email:          email,
		PasswordHashed: hashedPassword,
		Role:           "admin",
	}

	err := myValidator.Struct(newAdmin)

	if err != nil {
		return nil, err
	}

	return &newAdmin, nil
}

func (adm *Admin) ToJson() (string, error) {
	jsonBytes, err := json.Marshal(adm)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (adm *Admin) GetProperty(property string) (Serializeable, error) {

	var value MyString;

	switch property {
	case "email":
		value.CopyString(adm.Email)
	case "role":
		value.CopyString(adm.Role)
	case "password":
		value.CopyString(adm.PasswordHashed)
	default:
		return nil, fmt.Errorf("The property %s is not valid for an admin", property)
	}
	return &value, nil
}

func (adm *Admin) Validate() error {
	return myValidator.Struct(adm)
}

func (adm *Admin) SetProperty(key string, value Serializeable) error {
	return fmt.Errorf("Cannot Change Properties of an Admin")
}

func (adm *Admin) String() string {
	return fmt.Sprintf("User is an Admin | Email: %s |", adm.Email)
}
