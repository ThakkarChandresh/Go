package main

import (
	"fmt"
	"strings"
)

func main() {
	//we can only compare maps to nil we can't directly compare one map to another using ==
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No Of Employees: %d\n", len(employees))

	//we can get the value of the corresponding key even if the variable is not initialized!
	fmt.Printf("The value for key %q is %q\n", "Chandresh", employees["chandresh"])

	//we can use only comparable types for the keys so we cannot define slice as a key but we can define an array as a map key
	//var m1 map[[]int]string
	var m1 map[[2]int]string
	_ = m1

	//in go it's illegal to populate an un initialized map. We have to initialize the map using the make function or a map literal before we can add any elements.
	//employees["Chandresh"] = "Development" //assignment to entry in nil map

	//Way :1
	salary := map[string]int{}

	salary["Chandresh"] = 1_00_000
	salary["Nikunj"] = 1_00_102

	fmt.Println(salary)

	//Way :2
	departments := make(map[string]string)
	departments["Chandresh"] = "Development"
	departments["Nikunj"] = "HR"

	fmt.Println(departments)

	//Way :3
	techStack := map[string]string{
		"Chandresh": "Spring",
		"Nikunj":    "Figma", //this , is compulsory for multi line declaration but not for single line declaration
		//"Nikunj": "HJKK" //at the time of declaration the duplicate key is not allowed i.e it will be not overriden at the time of declaration.
	}

	//techStack["Nikunj"] = "HK" //but we can the value over some key like this.

	fmt.Println(techStack)

	m := map[int]int{1: 2, 3: 4, 5: 4}
	_ = m

	//Working with if key exists in map or not!
	fmt.Println(salary["Tiwari"]) //we got zero value but who knows if we added this entry with zero value!!

	salary["deven"] = 0

	//v, ok := salary["Tiwari"]
	v, ok := salary["deven"]

	if ok {
		fmt.Println("The Key Exists and value is", v)
	} else {
		fmt.Println("Key Not Exists and default zero value is", v)
	}

	//Iterate over a map
	for key, value := range salary {
		fmt.Printf("%s--------------%d\n", key, value)
	}

	fmt.Printf("\n%s\n", strings.Repeat("^", 10))

	delete(salary, "deven")
	for key, value := range salary {
		fmt.Printf("%s--------------%d\n", key, value)
	}
}
