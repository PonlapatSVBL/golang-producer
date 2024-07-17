package main

import (
	"fmt"

	"github.com/PonlapatSVBL/golang-producer/connections/mysql"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

func main() {
	utils.LoadEnv()

	type Employee struct {
		EmployeeId   string `json:"employee_id" db:"employee_id"`
		EmployeeName string `json:"employee_name" db:"employee_name"`
	}

	query := "SELECT employee_id, employee_name FROM hms_api.comp_employee LIMIT 1,2"

	mysqlInstance := mysql.NewMysql()
	employees := []Employee{}
	mysqlInstance.SqlList(&employees, query)

	utils.PrintJsonIndent(employees, true)
	fmt.Println()
}
