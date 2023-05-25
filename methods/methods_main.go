package main

import (
	"fmt"
	"strings"
)

// Method belongs to a type and function belongs to a package
type Friends []string

type myInt int

func (m myInt) increment() myInt {
	m += 1
	return m
}

func (f Friends) getFriends() {
	for k, v := range f {
		fmt.Println(k, v)
	}

	f[0] = "Heeyt"
}

func main() {
	myFriends := Friends{"Priyam", "Chittu", "Muskan", "Nikunj", "Dhaval"}

	myFriends.getFriends()

	fmt.Println(strings.Repeat("*", 20))

	// Method is ultimately is a function that takes a receiver as an argument
	Friends.getFriends(myFriends)

	fmt.Println(strings.Repeat("*", 20))

	var mi myInt = 100

	fmt.Println(mi)

	fmt.Println(mi.increment())

	fmt.Println(mi)
}
