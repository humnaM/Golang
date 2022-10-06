package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}
func print(stlist *StudentList) {

	var n = 0
	fmt.Printf("%s List %d%s\n", strings.Repeat("=", 25), n, strings.Repeat("=", 25))
	println("Name of the student is: ", stlist.list[0].name)
	println("The roll number:", stlist.list[0].rollnumber)
	println("Address:", stlist.list[0].address)

	n = 1
	fmt.Printf("%s List %d%s\n", strings.Repeat("=", 25), n, strings.Repeat("=", 25))
	println("Name of the student is: ", stlist.list[1].name)
	println("The roll number:", stlist.list[1].rollnumber)
	println("Address:", stlist.list[1].address)

}
func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")
	print(student)
}
