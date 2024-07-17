package service

import (
	"log"
	"reflect"

	"github.com/PonlapatSVBL/golang-producer/database/mysql"
	"github.com/PonlapatSVBL/golang-producer/models"
	"github.com/PonlapatSVBL/golang-producer/utils"
)

func GetEmployeeByServerId(serverId []string) ([]models.EmployeeStruct, error) {
	var employees []models.EmployeeStruct

	query := "SELECT employee_id, employee_name FROM hms_api.comp_employee LIMIT 1,2"

	rows, err := mysql.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// สร้าง slicer สำหรับเก็บค่าของคอลัมน์ที่ดึงมา
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Print(err)
			continue
		}

		// สร้าง struct ใหม่เพื่อเก็บข้อมูลที่ดึงมาจากฐานข้อมูล
		emp := models.EmployeeStruct{}

		// ใช้ reflection ในการ map ค่าจาก values เข้า struct
		structValue := reflect.ValueOf(&emp).Elem()
		for i, col := range columns {
			fieldName := utils.ToCamelCase(col) // ใช้ ToCamelCase ในการแปลงชื่อฟิลด์
			field := structValue.FieldByName(fieldName)
			if !field.IsValid() {
				continue
			}

			// ตรวจสอบว่า valuePtrs[i] ไม่ใช่ nil ก่อนที่จะทำการแปลง
			if valuePtrs[i] != nil {
				// ตรวจสอบประเภทของฟิลด์และทำการแปลงค่าให้เหมาะสม
				switch field.Kind() {
				case reflect.String:
					// แปลง []uint8 เป็น string
					if strVal, ok := values[i].([]uint8); ok {
						field.SetString(string(strVal))
					}
					// สามารถเพิ่ม case อื่นๆ ตามประเภทของฟิลด์ได้ตามต้องการ
				}
			}
		}

		// เพิ่ม struct ที่ map ได้เข้า slice employees
		employees = append(employees, emp)

		// พิมพ์ข้อมูล JSON ของ employee
		// utils.PrintJsonMarshal(emp)
	}

	return employees, nil
}
