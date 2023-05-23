package main

import "fmt"

type Teacher struct {
	id   int
	name string
	dep  string
}

type Student struct {
	roll    int
	name    string
	marks   []int
	teacher *Teacher
}

func (t Teacher) toString() string {
	return fmt.Sprintf("{Name: %s,Department: %s}", t.name, t.dep)
}

func (s Student) toString() string {
	return fmt.Sprintf("Name: %s\nRoll: %d\nTeacher: %s", s.name, s.roll, s.teacher.toString())
}

func main() {
	var teacher = Teacher{dep: "CSE", id: 1, name: "Meghna Mam"}
	var school = map[int]Student{}
	school[1] = Student{0, "Mihir", []int{10, 30}, &teacher}
	var stud = school[1]
	//changed the roll number
	stud.roll = 1

	//but above change is not reflected in original map
	fmt.Println("Roll no 0 changed to -->", school[1].roll)

	var XMenSchool = map[int]*Student{}
	XMenSchool[0] = &Student{
		roll:    10,
		name:    "Mahir",
		marks:   []int{100},
		teacher: &Teacher{1, "Prof. X", "Mutants"},
	}

	var studRef = XMenSchool[0]
	//value is change by pointer in value of map
	studRef.roll = 1 // equivalent: (*studRef).roll=1

	//change of roll is reflect in the original map
	//because it is changed by deferencing the pointer
	fmt.Println(XMenSchool[0])

	var anonStruct = struct {
		length  int
		breadth float64
	}{
		3,
		4.5,
	}

	fmt.Println(anonStruct)
}
