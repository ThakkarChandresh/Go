package main

import "fmt"

type student interface {
	getRollNo() int
}

type faculty interface {
	getSalary() int
}

type collage interface {
	student
	faculty
	getName() string
}

func info(c collage) {
	fmt.Println(c.getSalary(), c.getName(), c.getRollNo())
}

type sil struct {
	name           string
	rollno, salary int
}

func (s sil) getRollNo() int {
	return s.rollno
}

func (s sil) getSalary() int {
	return s.salary
}

func (s sil) getName() string {
	return s.name
}

func main() {
	myCol := sil{
		name:   "Silver Oak",
		rollno: 1001,
		salary: 10_00_000,
	}

	info(myCol) //if we comment line 36-38 the it will raise error
}
