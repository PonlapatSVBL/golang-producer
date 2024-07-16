package main

import (
	"log"

	service "github.com/PonlapatSVBL/golang-producer/services"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

func main() {
	utils.LoadEnv()
	utils.ConnectMysql()

	employees, err := service.GetEmployeeByServerId([]string{"1", "2", "3"})
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintFilteredJsonMarshal(employees)
}
