package main

import (
	"fmt"
)

type humna struct {
	nameof string
	ageof  int
}

func printhumna(hum humna) {
	fmt.Println("nameof", hum.nameof)
	fmt.Println("ageof", hum.ageof)
}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person
	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	var h humna
	h.nameof = "humna cutu"
	h.ageof = 22

	printhumna(h)
	//print values of the person
	printPerson(pers1)
	printPerson(pers2)
}

func printPerson(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age:", person.age)
	fmt.Println("Job: ", person.job)
	fmt.Println("Salary: ", person.salary)

}
