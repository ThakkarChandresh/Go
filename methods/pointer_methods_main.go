package main

import "fmt"

// Methods are only for value types not for method types
type Car struct {
	brand string
	price int
}

func updateCar(existing Car, newBrand string, price int) {
	existing.brand = newBrand
	existing.price = price
}

/*func (c Car) updateCarMethod(newBrand string, price int) {
	c.brand = newBrand
	c.price = price
}*/

func (c *Car) updateCarMethodPointer(newBrand string, price int) {
	c.brand = newBrand
	c.price = price
}

func main() {

	audiCar := Car{"Audi", 100000}

	updateCar(audiCar, "Renault", 200000)

	fmt.Println(audiCar)

	//audiCar.updateCarMethod("Renault", 200000)

	//(&audiCar).updateCarMethodPointer("Renault", 200000)
	audiCar.updateCarMethodPointer("Renault", 200000) //This is same as above compiler does this & thing for us in the backend

	fmt.Println(audiCar)

	var audiCarPtr *Car = &audiCar

	fmt.Println(*audiCarPtr)
	// Valid ways
	audiCarPtr.updateCarMethodPointer("VM", 10_000_0)
	fmt.Println(*audiCarPtr)

	(*audiCarPtr).updateCarMethodPointer("VM NEW", 20_000)
	fmt.Println(audiCarPtr)
	fmt.Println(*audiCarPtr)
}
