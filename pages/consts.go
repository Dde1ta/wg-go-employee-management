package pages

import (
	"fmt"
	"os"
	"path/filepath"

	"wg.dde1ta/services"
)

var dbfilepath string
var auth services.AuthService
var employee services.EmployeeService
var admin services.AdminService

func init() {
	value, exists := os.LookupEnv("DATAPATH")

	if exists {
		dbfilepath = filepath.Join(value + "/db.json")
	} else {
		dbfilepath = "data/db.json"
	}

	fmt.Println("Init: DBFILE set to", dbfilepath)

	auth = services.NewAuthService(dbfilepath)
	employee = services.NewEmployeeService(dbfilepath)
	admin = services.NewAdminService(dbfilepath)
}
