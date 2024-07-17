package main

import (
	employeeService "github.com/PonlapatSVBL/golang-producer/services/employee"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

func main() {
	utils.LoadEnv()

	employees, _ := employeeService.GetEmployeeByServerId([]string{"1", "2"})
	utils.PrintExistJsonIndent(employees, true)

	// utils.ExportModel("hms_api.comp_employee")
}
