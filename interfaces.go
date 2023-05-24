package main

import "fmt"

type Food interface {
	PrintName(name string)
}

type Fruit string

type Drink string

func (d Drink) PrintName(name string) {
	fmt.Println("Drink Name : ", name)
}

func (f Fruit) PrintName(name string) {
	fmt.Println("Fruit Name : ", name)
}

func main() {
	var fruit Food = Fruit("")

	var drink Food = Drink("")

	fruit.PrintName("Apple")

	drink.PrintName("Pepsi")

}
