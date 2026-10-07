package entity

import "github.com/go-playground/validator"

var myValidator *validator.Validate

func init() {
	myValidator = validator.New()
}

func GetValidator() *validator.Validate {
	return myValidator
}

// package entity

// import (
// 	"github.com/go-playground/validator"
// )

// var creatorFunction = validatorStore()

// func validatorStore() func() *validator.Validate {
// 	var myValidator *validator.Validate = nil

// 	var createValidator = func () *validator.Validate {
// 		if myValidator == nil {
// 			myValidator = validator.New()
// 		}
// 		return myValidator
// 	}

// 	return createValidator

// }

// func GetValidator() *validator.Validate {
// 	return creatorFunction()
// }
