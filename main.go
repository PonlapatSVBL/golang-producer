package main

import (
	"log"

	"github.com/PonlapatSVBL/golang-producer/database/mysql"
	service "github.com/PonlapatSVBL/golang-producer/services"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

func main() {
	utils.LoadEnv()
	mysql.Connect()

	employees, err := service.GetEmployeeByServerId([]string{"1", "2", "3"})
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintFilteredJsonMarshal(employees)
}
