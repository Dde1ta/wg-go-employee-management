# Employee Management


## JSON Structure

This is the initial stucture for a simple employee management system json file.
```json
// db.json
[
    {
        "id": "string",
        "name": "string",
        "email": "string",
        "password_hashed": "string",
        "contact_number": "int",
        "department": "string",
        "position": "string",
        "role": "employee"
    },
    {
        "role": "admin",
        "id": "string",
        "email": "string",
        "password_hashed": "string"
    },
]
// logs.json

[
    {
        "session": "string",
        "principal": {
            "role": "admin | employee",
            "id" : "string"
        },
        "logged_in":  "timestamp",        
        "actions": [
            {
                "effect": "LOGIN | SIGUP | UPDATE | CREATE | DELETE",
                "resource": "str",
                "acted_at": "timestamp"                    
            }
        ]
    }
]

```

## Structs

I would need a lot of structs.

1. `FileManager`: Used to read and write to the file. Takes lock from `FileLock`, Writes `DB`
  2. `FileLock`: Using flock mentain the lock of the file, give lock to `FileManager`
3. `Employee`: Represents an employee's data
4. `Admin`: Represents an admin's data
5. `Log`: Represents a logs
6. `DB`: Mentain the JSON data -> uses the structs `Employee`, `Admin`, `Log` for it
7. `Logger`: Adds `Log` To `DB`

8. `AdminService`: Interact with `DB` to fulfill `Admin` actions
9. `EmployeeService`: Interact with `DB` to fulfill `Employee` actions
10. `AuthService`: Allow Login and Signup

11. `Context`: Stores all the info of the user
## Functions

I would need Functions which will run as scripts acting like the frontend which the user will interact with, It will be called by `main`.


3. AdminScript
4. EmployeeScript



# PassWord Reading


```go

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"syscall"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {

	var password string

	fmt.Print("Enter Password: ")

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("\nError reading password: %v\n", err)
		return
	}

	fmt.Println() 

	password = string(bytePassword)

	fmt.Println("The Captured Password:", password)

	hash, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hashed Password:", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Password Match Status:", match)
}


```