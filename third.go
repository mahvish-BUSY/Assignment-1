package main

import (
	"fmt"
	"reflect"
)

func populateStruct(source map[string]interface{}, structPtr interface{}) error {

	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		return fmt.Errorf("structPtr must be a pointer to a struct")
	}

	//retrieving the value structPtr points to
	structVal := reflect.ValueOf(structPtr).Elem()

	//iterating over the source map

	for key, value := range source {

		//for this key checking if it is present as a field in structVal
		field := structVal.FieldByName(key)

		//if field is present in struct
		if field.IsValid() {
			//if field is of kind struct
			if field.Kind() == reflect.Struct {
				//if corresponding value is a map
				if nestedMap, ok := value.(map[string]interface{}); ok {
					//create a new ptr to instance of field's type
					nestedStruct := reflect.New(field.Type()).Interface()
					if err := populateStruct(nestedMap, nestedStruct); err != nil {
						return err
					}

					//set this struct's value to the field
					field.Set(reflect.ValueOf(nestedStruct).Elem())
				}
			} else {
				field.Set(reflect.ValueOf(value))
			}

		} else if reflect.ValueOf(value).Kind() == reflect.Map {

			//if the coorresponding value of key is a map[string]interface{}
			if nestedMap, ok := value.(map[string]interface{}); ok {
				if err := populateStruct(nestedMap, structPtr); err != nil {
					return err
				}
			}

		} else if reflect.ValueOf(value).Kind() == reflect.Slice {
			for _, sliceVal := range value.([]interface{}) {
				if nestedMapInSlice, ok := sliceVal.(map[string]interface{}); ok {
					if err := populateStruct(nestedMapInSlice, structPtr); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

type Address struct {
	City    string
	Country string
}

type Education struct {
	Course      string
	PassingYear int
	Percentage  float64
}
type Person struct {
	Name            string
	Age             int
	Current_Address Address
}

func main() {

	myData := map[string]interface{}{
		"Name": "Muskan",
		"Age":  23,
		"Current_Address": map[string]interface{}{
			"City":    "Delhi",
			"Country": "India",
		},
		"permanent address": map[string]interface{}{
			"City":    "Kanpur",
			"Country": "India",
		},
		"education": []interface{}{
			map[string]interface{}{
				"Course":      "B.Sc",
				"PassingYear": 2021,
				"Percentage":  78.88,
			},
			map[string]interface{}{
				"Course":      "MCA",
				"PassingYear": 2024,
				"Percentage":  86.7,
			},
		},
		"intern":  true,
		"company": "BUSY Infotech Pvt Ltd",
		"CompanyAddress": map[string]interface{}{
			"City":    "Delhi",
			"Country": "India",
		},
		"mentorAssigned": []interface{}{"A", "B", "C"},
	}

	var addressPtr = &Address{}
	var educationPtr = &Education{}
	var personPtr = &Person{}

	//populating struct instance which have a field which is also a struct variable
	if err := populateStruct(myData, personPtr); err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Printf("%+v\n", personPtr)
	}

	//populating struct instance present in nested map
	if err := populateStruct(myData, addressPtr); err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Printf("%+v\n", addressPtr)
	}

	//populating struct instance present in []interface{}
	if err := populateStruct(myData, educationPtr); err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Printf("%+v\n", educationPtr)
	}

}
