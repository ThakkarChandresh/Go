package main

import "fmt"

func Print(s []int) {
	fmt.Println("length:", len(s), "\tcapacity: ", cap(s), "\telements: ", s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13, 2, 3, 5, 7, 11, 13}
	Print(s)

	/*s1 := s[2:] //[]
	Print(s1)

	s2 := s1[:6] //[2, 3, 5, 7]
	Print(s2)

	s3 := s2[0:] //[5, 7]
	Print(s3)

	s := s3[0:12]
	Print(s)*/

	s = s[2:] //[]
	Print(s)

	s = s[:6] //[2, 3, 5, 7]
	Print(s)

	s = s[3:] //[5, 7]
	Print(s)

	s = s[0:7]
	Print(s)

}
