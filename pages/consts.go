package pages

import (
	"wg.dde1ta/services"
)

const DBFILEPATH string = "data/db.json"

var auth services.AuthService = services.NewAuthService(DBFILEPATH)

var employee services.EmployeeService = services.NewEmployeeService(DBFILEPATH)

var admin services.AdminService = services.NewAdminService(DBFILEPATH)