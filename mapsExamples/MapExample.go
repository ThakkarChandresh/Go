package main

import "fmt"

func main() {

	var someMap map[string]int

	//someMap["delhi"] = 110000 // Assignment to nil map is not allowed

	fmt.Println(someMap)

	var harshMap = map[string]map[string]int{"ahmedabad": {"gota": 380060, "sola": 380059}}

	fmt.Println(harshMap)

	harshMap["gandhinagar"] = map[string]int{"adalaj": 382424, "chandkheda": 382421}

	fmt.Println(harshMap)

	fmt.Println(harshMap["ahmedabad"])

	fmt.Println(harshMap["ahmedabad"]["gota"])

	if elem, ok := harshMap["gandhinagar"]; ok {

		fmt.Println("value for key: gandhinagar found: ", elem)

		if elem1, ok1 := elem["adalaj"]; ok1 {

			if ok1 {
				fmt.Println("value for key adalaj found: ", elem1)
			} else {
				fmt.Println("value for key: adalaj not found")
			}
		}
	} else {
		fmt.Println("value for key: gandhinagar not found")
	}

}
