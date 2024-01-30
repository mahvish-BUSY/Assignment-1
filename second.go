package main

import (
	"fmt"
	"reflect"
)

func removeKey(key string, source map[string]interface{}) (map[string]interface{}, error) {
	sourceVal := reflect.ValueOf(source)

	switch sourceVal.Kind() {
	case reflect.Map:

		//checking if key is present in current map
		//It returns the zero Value if key is not found in the map or if v represents a nil map.
		val := sourceVal.MapIndex(reflect.ValueOf(key))

		//checking if val is a valid value
		if val.IsValid() {
			//remove the key using func (Value) SetMapIndex
			sourceVal.SetMapIndex(reflect.ValueOf(key), reflect.Value{})
			fmt.Println("Key and value pair removed successfully !!")
			return source, nil
		}

		//looking for nested map or []interface{} in values of current map

		//extracting keys of current map
		keys := sourceVal.MapKeys() //keys is []Value

		//iterating over this keys slice and looking for values if it is a map[string]interface{}

		for _, keyInMap := range keys {
			//extracting value from sourceVal for this particular key
			mapVal := sourceVal.MapIndex(keyInMap) //key is already of type 'Value'

			//checking if it is a map[string]interface{}
			if nestedMap, ok := mapVal.Interface().(map[string]interface{}); ok {

				if returnedMap, err := removeKey(key, nestedMap); err == nil {
					return returnedMap, nil
				}
			}

			// checking the map val if it is []interface{}
			if nestedSlice, ok := mapVal.Interface().([]interface{}); ok {
				for _, sliceVal := range nestedSlice {

					if nestedMap, ok := reflect.ValueOf(sliceVal).Interface().(map[string]interface{}); ok {
						if returnedMap, err := removeKey(key, nestedMap); err == nil {
							return returnedMap, nil
						}
					}
				}
			}

		} //for loop terminates

	default:
		return source, fmt.Errorf("no such type of key")

	}

	return source, fmt.Errorf("key is not present")
}
func main() {
	myData := map[string]interface{}{
		"name": "Muskan",
		"age":  23,
		"info": map[string]interface{}{
			"city":    "Delhi",
			"pincode": 110025,
			"country": "India",
			"permanentAdd": map[string]interface{}{
				"city":     "kanpur",
				"country":  "India",
				"pincode":  208004,
				"landmark": "Pandit Hotel",
				"state":    "Uttar Pradesh",
			},
			"mobile": 123456789,
		},
		"education": []interface{}{
			map[string]interface{}{
				"course":      "B.Sc",
				"passingYear": 2021,
				"percentage":  78.88,
			},
			map[string]interface{}{
				"course":      "MCA",
				"passingYear": 2024,
				"CGPA":        8.02,
			},
		},
		"intern":         true,
		"company":        "BUSY Infotech Pvt Ltd",
		"mentorAssigned": []interface{}{"A", "B", "C"},
	}

	key := "percentage"
	// var value interface{}
	// value = "Delhi"
	fmt.Println("Map before modification : ", myData)
	if _, err := removeKey(key, myData); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myData)
	}
}
