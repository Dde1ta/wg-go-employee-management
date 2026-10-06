package pages

import (
	"fmt"
	"os"

	"wg.dde1ta/services"
)

var DBFILEPATH string
var auth services.AuthService
var employee services.EmployeeService
var admin services.AdminService

func init() {
	value, exists := os.LookupEnv("DBFILEPATH")

	if exists {
		DBFILEPATH = value
	} else {
		DBFILEPATH = "data/db.json"
	}

	fmt.Println("Init: DBFILE set to", DBFILEPATH)

	auth = services.NewAuthService(DBFILEPATH)
	employee = services.NewEmployeeService(DBFILEPATH)
	admin = services.NewAdminService(DBFILEPATH)
}
