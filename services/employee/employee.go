package employeeService

import (
	"github.com/PonlapatSVBL/golang-producer/connections/mysql"
	"github.com/PonlapatSVBL/golang-producer/models"
)

func GetEmployeeByServerId(serverId []string) ([]models.EmployeeStruct, error) {
	var employees []models.EmployeeStruct

	query := "SELECT employee_id, employee_name FROM hms_api.comp_employee LIMIT 1,2"

	mysqlInstance := mysql.NewMysql()

	employees = []models.EmployeeStruct{}
	mysqlInstance.SqlList(&employees, query)

	return employees, nil
}
