package main

import (
	"fmt"
	"reflect"
)

func setKeyValue(key string, value interface{}, source map[string]interface{}) (map[string]interface{}, error) {
	return setKeyValueRecursive(key, value, source)
}
func setKeyValueRecursive(key string, value interface{}, source map[string]interface{}) (map[string]interface{}, error) {

	sourceValue := reflect.ValueOf(source)

	switch sourceValue.Kind() {

	case reflect.Map:

		//checking if the key exists in current map if yes set the value
		val := sourceValue.MapIndex(reflect.ValueOf(key))
		//checking if the value returned by MaPIndex() is valid i.e not a zero(reflect.Zero) if valid then key is present in map
		if val.IsValid() {
			sourceValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
			fmt.Println("Value updated successfully!! ")

			return source, nil
		}

		//retrieving all keys from current map
		keys := sourceValue.MapKeys() //MapKeys returns a slice containing all the keys present in the map, in unspecified order.

		//iterating over this slice of keys and checking whether any of the keys value is a map[strig]interface{} or a []interface{}
		for _, mapKey := range keys {

			//retrieving associated value to the key in reflect.Value type
			mapVal := sourceValue.MapIndex(mapKey)

			if nestedMap, ok := mapVal.Interface().(map[string]interface{}); ok {
				if nestedMap, err := setKeyValueRecursive(key, value, nestedMap); err == nil {
					return nestedMap, nil
				}
			}

			// checking the map val if it is []interface{}
			if nestedSlice, ok := mapVal.Interface().([]interface{}); ok {
				for _, sliceVal := range nestedSlice {

					if nestedMap, ok := reflect.ValueOf(sliceVal).Interface().(map[string]interface{}); ok {
						if returnedMap, err := setKeyValueRecursive(key, value, nestedMap); err == nil {
							return returnedMap, nil
						}
					}
				}
			}

		}

	default:
		return source, fmt.Errorf("no Such type of key")

	}

	//key notfound

	return source, fmt.Errorf("key not found: %s", key)
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
			"mobile": 9453660133,
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

	key := "passingYear"
	var value interface{}
	value = 2022
	if _, err := setKeyValue(key, value, myData); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myData)
	}
}
