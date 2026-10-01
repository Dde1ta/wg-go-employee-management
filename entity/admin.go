package entity

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Admin struct {
	Id string             `json:"id" validate:"required"`
	Email string 		  `json:"email" validate:"required,email"`
	PasswordHashed string `json:"password" validate:"required"`
}

func NewAdmin(email, password string) (Admin, error) {
	var newId, err = uuid.NewV7()

	if err != nil {
		return Admin{}, err
	}

	newAdmin := Admin{
		Id: newId.String(),
		Email: email,
		PasswordHashed: password,
	}

	err = myValidator.Struct(newAdmin)

	if err != nil {
		return Admin{}, err
	}
	
	return newAdmin, nil
}

func (adm Admin) ToJson() (string, error){
	jsonBytes, err := json.Marshal(adm)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (adm Admin) GetId() (string) {
	return adm.Id
}
