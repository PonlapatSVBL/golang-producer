package utils

import (
	"encoding/json"
	"log"
	"reflect"
)

func PrintJsonMarshal(data interface{}) {
	// แปลง slice ของ employees เป็น JSON
	empJSON, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling employees to JSON:", err)
	}

	// พิมพ์ JSON ออกมา
	log.Println(string(empJSON)) // หรือใช้ fmt.Println(string(empJSON)) ตามที่ต้องการ
}

func PrintFilteredJsonMarshal(data interface{}) {
	// ใช้ reflection เพื่อตรวจสอบ key ที่มีค่าว่างและลบออก
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		log.Println("Input data must be a slice")
		return
	}

	// สร้าง slice ใหม่เพื่อเก็บข้อมูลที่มีค่าเท่ากับศูนย์
	var filteredSlice []interface{}
	for i := 0; i < value.Len(); i++ {
		itemValue := value.Index(i)
		if itemValue.Kind() == reflect.Struct {
			newItem := make(map[string]interface{})
			for j := 0; j < itemValue.NumField(); j++ {
				field := itemValue.Type().Field(j).Name
				fieldValue := itemValue.Field(j)
				if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
					continue
				}
				newItem[field] = fieldValue.Interface()
			}
			filteredSlice = append(filteredSlice, newItem)
		} else {
			log.Println("Non-struct item found in slice")
			continue
		}
	}

	// แปลง slice ใหม่เป็น JSON
	empJSON, err := json.Marshal(filteredSlice)
	if err != nil {
		log.Println("Error marshalling data to JSON:", err)
		return
	}

	// พิมพ์ JSON ออกมา
	log.Println(string(empJSON)) // หรือใช้ fmt.Println(string(empJSON)) ตามที่ต้องการ
}
