package main

import "fmt"

func main() {
	type Man struct {
		Name     string
		LastName string
		Age      int
		Gender   string
		Crimes   int
	}

	people := map[string]Man{
		"Barry":  {Name: "Barry", LastName: "Qup", Age: 41, Gender: "male", Crimes: 2},
		"Ruby":   {Name: "Ruby", LastName: "Qaz", Age: 19, Gender: "male", Crimes: 9},
		"Andy":   {Name: "Andy", LastName: "Indry", Age: 33, Gender: "male", Crimes: 0},
		"Ingrid": {Name: "Ingrid", LastName: "Beetle", Age: 35, Gender: "female", Crimes: 2},
		"Shoot":  {Name: "Shoot", LastName: "First", Age: 44, Gender: "male", Crimes: 4},
	}
	suspects := []string{"Barry", "Andy", "Ingrid", "Shoot"}
	var moreCriminal Man
	var CriminalSearch bool
	for _, name := range suspects {
		citizen, ok := people[name]
		if !ok {
			continue
		}
		if citizen.Crimes > moreCriminal.Crimes {
			moreCriminal = citizen
			CriminalSearch = true
		}
	}
	if CriminalSearch {
		fmt.Printf("Самый криминальный: %s %s\n", moreCriminal.Name, moreCriminal.LastName)
	} else {
		fmt.Println("Запрошенный люди в базе не числятся")
	}

}
