package main

import "fmt"

func printProduct(x, y int) {
	fmt.Println(x * y)
}

// returns multiple values
func getMultipleValue() (int, string, float64) {
	return 10, "mihir", 45.45
}

func sum(a, b int) any {
	defer func() {
		a = a + b
		fmt.Printf("%p\n", &a)
	}()
	fmt.Printf("%p\n", &a)
	return a
}

// !!!Below will give error as function with name "sum" is already declared
//func sum(a, b, c int) {
//	defer func() {
//		a = a + b
//		fmt.Printf("%p\n", &a)
//	}()
//	fmt.Printf("%p\n", &a)
//}

func update(a *int) {
	*a = 314 // deferencing pointer address
	return
}

func Do(a int) {

	//define anonymous function inside function
	func() {
		a++
	}()

	func() {
		a++
	}()

	//define struct inside function
	type Car struct {
		tyres int
		price float64
	}

	var carName Car = Car{4, 2.1}
	fmt.Println(carName)

	fmt.Println(a)
}

// changing array inside function don't
// affect the original array
func changeArray(arr [3]int) {
	arr[0] = 1
	fmt.Println("arr: ", arr)
}

func driveCar(drive *func(string), speed float64) {
	(*drive)(fmt.Sprintf("Car is driving at %f", speed))
}

func driveCar1(drive func(string), speed float64) {
	drive(fmt.Sprintf("Car is driving at %f", speed))
}

func main() {

	var fun = func(s string) {
		fmt.Println(s)
	}
	driveCar(&fun, 100.1)
	driveCar1(fun, 102.13)

	Do(5)
	fmt.Println("sum: ", sum(10, 20))
	printProduct(10, 20)
	fmt.Println(getMultipleValue())

	var arr = [3]int{1, 2, 3}
	fmt.Println("Before:", arr)
	update(&arr[0])
	fmt.Println("After :", arr)

	changeArray(arr)
	fmt.Println("arr: ", arr)

	//assigning anon functions to variables
	var draw = func(painting string) {
		fmt.Println("Drawing", painting)
	}
	draw("dog")

}
