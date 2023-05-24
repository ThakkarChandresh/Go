package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

type Element interface{}

type Stack struct {
	elements []Element
	mutex    *sync.Mutex
}

func newEmptyStack() *Stack {
	return &Stack{
		elements: nil,
	}
}

func newStack(elements []Element, mutex *sync.Mutex) *Stack {
	return &Stack{
		elements: elements,
		mutex:    mutex,
	}
}

func (stack *Stack) Push(element Element) {

	stack.mutex.Lock()

	defer stack.mutex.Unlock()

	stack.elements = append(stack.elements, element)

}

func (stack *Stack) Pop() (element Element) {

	stack.mutex.Lock()

	defer stack.mutex.Unlock()

	if len(stack.elements) == 0 {
		return nil
	}

	element = stack.elements[len(stack.elements)-1]

	stack.elements = stack.elements[:len(stack.elements)-1]

	return
}

func (stack *Stack) Top() (element Element) {
	stack.mutex.Lock()

	defer stack.mutex.Unlock()

	if len(stack.elements) == 0 {
		return nil
	}

	element = stack.elements[len(stack.elements)-1]

	return
}

func (stack *Stack) isEmpty() (result bool) {

	stack.mutex.Lock()

	defer stack.mutex.Unlock()

	result = len(stack.elements) == 0

	return
}

func (stack *Stack) Size() (size int) {

	stack.mutex.Lock()

	defer stack.mutex.Unlock()

	size = len(stack.elements)

	return
}

func (stack *Stack) Clear() {

	stack.mutex.Lock()

	defer stack.mutex.Unlock()

	stack.elements = []Element{}
}

func task(stack *Stack) {

	defer wg.Done()

	stack.Pop()
	stack.Pop()
	stack.Pop()
}

func task2(stack *Stack) {

	defer wg.Done()

	stack.Push(60)
	//fmt.Printf("%#v\n", stack)
	stack.Push(70)
	//fmt.Printf("%#v\n", stack)
	//stack.Push(80)
	//fmt.Printf("%#v\n", stack)
}

func main() {

	wg.Add(2)

	elements := []Element{10, 20, 30, 40, 50}

	var mutex = sync.Mutex{}

	var stack = newStack(elements, &mutex)

	go task(stack)

	go task2(stack)

	wg.Wait()

	fmt.Println(stack.elements)

	//fmt.Println(stack.elements)

	//fmt.Println("Size of Stack : ", stack.Size())
	//
	//fmt.Println("Top of Stack : ", stack.Top())
	//
	//fmt.Println("Pushing element 60 in Stack...")
	//
	//stack.Push(Element(60))
	//
	//fmt.Println("After pushing size of Stack : ", stack.Size())
	//
	//fmt.Println("After pushing top of Stack : ", stack.Top())
	//
	//fmt.Println("Now Popping three times from stack...")
	//
	//stack.Pop()
	//
	//stack.Pop()
	//
	//stack.Pop()
	//
	//fmt.Println("Size of Stack after pop operations : ", stack.Size())
	//
	//fmt.Println("Top of Stack after pop operations : ", stack.Top())
	//
	//fmt.Println("Now clearing stack...")
	//
	//stack.Clear()
	//
	//fmt.Println("Size of Stack after clearing : ", stack.Size())

}
