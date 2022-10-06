package main

import (
	"fmt"
)

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	Alice := employee{"Alice", 10000, "Full-Stack Developer"}
	Lucy := employee{"Lucy", 7000, "Back-end Developer"}
	Charles := employee{"Charles", 4000, "Front-end Developer"}

	employees := []employee{Alice, Lucy, Charles}
	company := company{"PTCL", employees}

	fmt.Printf("Company is %v", company)
}
