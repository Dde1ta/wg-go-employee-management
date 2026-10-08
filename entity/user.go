package entity

import "fmt"

type MyString string

func (s *MyString) ToString() string {
	return string(*s)
}

func (s *MyString) CopyString(newS string) {
	*s = MyString(newS)
}

type Serializeable interface {
	ToString() string
	CopyString(string)
}

type User interface {
	ToJson() (string, error)
	GetProperty(string) (Serializeable, error)
	Validate() error
	SetProperty(string, Serializeable) error
	fmt.Stringer
}
