package main

import "fmt"

type Employees struct {
	string
	int
}

type Friends struct {
	name    string
	contact int
	age     int
	string
}

func main() {
	emp1 := Employees{"Chan", 11}
	fmt.Printf("%+v\n", emp1)
	fmt.Println(emp1.int, emp1.string)

	f1 := Friends{
		name:    "Nikunj",
		contact: 82982882,
		age:     12,
		string:  "GoodFriend",
	}

	fmt.Printf("\n%+v\n", f1)
	fmt.Println(f1.string, f1.name, f1.contact, f1.age)

}
