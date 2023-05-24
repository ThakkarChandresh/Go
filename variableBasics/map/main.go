package main

import "fmt"

func main() {
	var employees = map[string]int{"Mark": 10, "Sandy": 20}

	map1 := map[string]int{"Mark1": 10, "Sandy1": 20}

	fmt.Println(map1)

	fmt.Println(employees)

	map2 := make(map[string]int)

	fmt.Println(map2)

	map3 := make(map[string]int)

	var map4 = map[string]int{}

	fmt.Println("map3: ", len(map3), " | map4: ", len(map4))

	var harsh = map[string]int{}
	var harsh1 = map[string]int{}

	harsh["name"] = 10
	harsh1["name"] = 10
	//harsh["name"] = "kumar"

	s1 := fmt.Sprintf("%s", harsh)
	s2 := fmt.Sprintf("%s", harsh1)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Both map are same")
	} else {
		fmt.Println("Both map are not same")
	}

	map5 := map4

	map5["Nikunj"] = 100

	fmt.Println(map4, map5)

	delete(map5, "Nikuj")
}
