package main

func main() {
	sendOnlyChannel := make(<-chan int)

	readOnlyChannel := make(chan<- int)

	biDirectional := make(chan int)

	uniOne(sendOnlyChannel)
	//uniOne(readOnlyChannel) //error
	uniOne(biDirectional)

	//uniTwo(sendOnlyChannel) //error
	uniTwo(readOnlyChannel)
	uniTwo(biDirectional)

	//bi(sendOnlyChannel) //error
	//bi(readOnlyChannel) //error
	bi(biDirectional)
}

func uniOne(uniChannel <-chan int) {

}
func uniTwo(uniChannel chan<- int) {

}

func bi(biChannel chan int) {

}
