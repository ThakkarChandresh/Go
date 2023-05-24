package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
		defer println("defer: ", i)
	}
}
