package main

import (
	"fmt"
)

type Job struct {
	Title  string
	Salary int
}

func (job Job) displayJobInformation() string {
	return fmt.Sprintf("%s - £%d", job.Title, job.Salary)
}

func main() {
	softwareDeveloper := Job{"Software Developer", 25000}
	fmt.Println(softwareDeveloper.displayJobInformation()) // Software Developer (£25000)
}
