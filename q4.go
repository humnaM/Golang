package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type student struct {
	rollnumber int
	name       string
	address    string
	subjects   [3]string
	hash       string
}

func CalculateHash(hashh string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func newstudent(rollno int, name string, address string, subject [3]string) *student {
	s := new(student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subjects = subject
	subject2 := subject[0] + subject[1] + subject[2]
	s.hash = CalculateHash(strconv.Itoa(rollno) + name + address + subject2)

	return s
}

type studentlist struct {
	list []*student
}

func (obj *studentlist) createstudent(rollno int, name string, address string, subject [3]string) *student {

	as1 := newstudent(rollno, name, address, subject)
	obj.list = append(obj.list, as1)
	return as1
}

func printstudent(obj *studentlist, size int) {

	for i := 0; i < size; i++ {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))

		fmt.Println("Roll no: ", obj.list[i].rollnumber)
		fmt.Println("Name: ", obj.list[i].name)
		fmt.Println("Address: ", obj.list[i].address)
		fmt.Println("Subjects: ", obj.list[i].subjects)
		fmt.Println("Hash: ", obj.list[i].hash)
	}

}

func main() {

	s1 := new(studentlist)
	var subject1 = [3]string{"PDC", "SSD", "Blockchain"}
	var subject2 = [3]string{"PF", "CNet", "Entrapreneurship"}
	s1.createstudent(1, "Humna", "Islamabad", subject1)
	s1.createstudent(2, "Mursalin", "Islambad", subject2)

	printstudent(s1, 2)
}
