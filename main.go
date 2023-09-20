package main

import (
	"fmt"
	"ngc5/hero"
	"ngc5/person"
)

func main() {
	// ------------ NGC-5 : Struct and Method 1 ----------------
	p := person.Person{
		Name: "John",
		Age:  45,
		Job:  "Gambler",
	}

	p.GetInfo()

	// Add Year sebanyak 5x
	p.AddYear()
	p.AddYear()
	p.AddYear()
	p.AddYear()
	p.AddYear()

	fmt.Println("----- Setelah Add Year -----") // Separator

	p.GetInfo() // Get info setelah add year 5x
	// ---------------------------------------------------------

	// --------------- NGC-5 : Struct and Method 2 ---------------
	data := []map[string]any{
		{
			"name": "Bambang",
			"age":  20,
			"job":  "Gambler",
		},
		{
			"name": "Saputra",
			"age":  30,
			"job":  "Driver",
		},
		{
			"name": "John",
			"age":  40,
			"job":  "Police",
		},
		{
			"name": "Doe",
			"age":  45,
			"job":  "Programmer",
		},
	}

	for _, value := range data {
		name := value["name"].(string)
		age := value["age"].(int)
		job := value["job"].(string)

		v := person.Person{
			Name: name,
			Age:  age,
			Job:  job,
		}

		fmt.Println("------------ Person Loop ------------") // Separator
		v.GetInfo()
		fmt.Println("-------------------------------------") // Separator
	}
	// ---------------------------------------------------------------

	// --------------- NGC-5 : Struct Hero -----------------
	penyerang := map[string]int{
		"name":           1,
		"baseAttack":     10,
		"defence":        10,
		"criticalDamage": 10,
		"healthPoint":    50,
		"weapon":         10,
	}

	diserang := map[string]int{
		"name":           2,
		"baseAttack":     10,
		"defence":        10,
		"criticalDamage": 10,
		"healthPoint":    50,
		"weapon":         10,
	}

	fmt.Println()
	fmt.Println("------------- Hero ------------------")
	hero.Battle(penyerang, diserang)
	// -----------------------------------------------------
}
