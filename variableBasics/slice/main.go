package main

import "fmt"

func main() {
	var a = make([]int, 1, 10)

	a[0] = 10
	c := append(a, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89)

	fmt.Println("c: ", c)

	fmt.Println("a: ", a)
	fmt.Println("Capacity: ", cap(a), " | Length: ", len(a))

	a = append(a, 99, 98, 97, 96, 95, 94, 93, 92, 91)
	fmt.Println("Capacity: ", cap(a), " | Length: ", len(a))
	var b = new([40]int)[5:10]

	b = append(b, 5, 4, 3, 2, 1)

	fmt.Println("b: ", b)

	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 3)
	fmt.Println(strSlice)

	d := make([]int, 2, 5)

	d[0] = -1

	d[1] = -2

	copy(d, a)

	fmt.Println("inside d: ", d)

	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}

	slice2 = append(slice2, slice1...)

	fmt.Println("slice2: ", slice2)

	var slice3 []int

	fmt.Println(slice3, " comparision: ", slice3 == nil)

	var slice4 = []int{}

	fmt.Println(slice4, " comparision: ", slice4 == nil)

	var slice5 = make([]int, 0, 5)

	fmt.Println(slice5 == nil)

	a = make([]int, 0, 5)

	a = c

	fmt.Println("a: ", a)

	c[1] = -10

	fmt.Println("a: ", a)

	slice6 := make([]int, 10, 10)

	n := copy(slice6, c)

	fmt.Println("n: ", n, "Slice6: ", slice6)

	c[1] = -200

	fmt.Println("Slice6: ", slice6)

	slice7 := []string{"a", "b", "c", "d", "e", "f"}
	slice8 := append(slice7[0:1], "x", "y")

	fmt.Println(slice8, len(slice8), cap(slice8))

	fmt.Println("C:\\Users\\Andrei")

	s1 := "Golang"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v", s1[i])
	}

	fmt.Println(s1[0:1])

	s2 := s1
	_ = s2
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
