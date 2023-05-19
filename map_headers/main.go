package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	//when declaring a map variable go creates a map header in the memory
	//the map references this internal data structure, the map header
	//the map contains only the address of map header the key value paris of the map are not directly stored into the map
	//they are stored in the memory address referenced by the map header

	//so when you copy a map to a new map internal data structure is not copied, just referenced.
	friends := map[string]string{
		"chandresh": "deven",
		"deven":     "nikunj",
		"nikunj":    "chandresh",
	}

	contacts := friends
	contacts["deven"] = "chandresh"

	fmt.Println(contacts)
	fmt.Println(friends)

	newFriends := make(map[string]string)

	for key, value := range friends {
		newFriends[key] = value
	}

	newFriends["deven"] = "nikunj"
	fmt.Printf("\n%s\n", strings.Repeat("^", 10))

	fmt.Println(contacts)
	fmt.Println(friends)
	fmt.Println(newFriends)

	delete(newFriends, "ABC") //it will not generate any error even if key not exists.

	newFriends["deven"] = "chandresh"
	fmt.Println(reflect.DeepEqual(newFriends, friends))

	//dummyFriends := make(map[string]string)
	//reflect.Copy(dummyFriends, newFriends)

}
