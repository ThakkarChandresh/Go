package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	var c chan string

	c = make(chan string)

	var m sync.Mutex

	go func() {
		m.Lock()

		c <- "Lock Acquired!"

		c <- "Lock Acquired!"
	}()

	go func() {

		time.Sleep(2 * time.Second)

		res := <-c

		fmt.Println("Received By GR1", res)

		fmt.Println("I'm Trying to get lock", m.TryLock())

		fmt.Println("Let's try unlock the lock acquired by GR1!")

		m.Unlock()

		fmt.Println("Again I'm Trying to get lock", m.TryLock())

		close(c)

	}()

	go func() {

		fmt.Printf("Active %d\n", runtime.NumGoroutine())

		res := <-c

		fmt.Println("Received By GR2", res)

		fmt.Println("I'm Trying to get lock", m.TryLock())

		fmt.Println("Let's try unlock the lock acquired by GR2!")

		m.Unlock()

		fmt.Println("Again I'm Trying to get lock", m.TryLock())

	}()

	time.Sleep(time.Second * 10)

	<-c
}
