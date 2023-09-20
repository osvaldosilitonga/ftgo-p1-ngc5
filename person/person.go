package person

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  string
}

func (person Person) GetInfo() {
	fmt.Println("Name :", person.Name)
	fmt.Println("Age :", person.Age)
	fmt.Println("Job :", person.Job)
}

func (person *Person) AddYear() {
	person.Age += 1

	if person.Age >= 50 {
		person.Job = "Retired"
	}
}
