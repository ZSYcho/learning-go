package main

import "fmt"

func main() {

	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Println(studentGrades)
	studentGrades["Alice"] = 95
	fmt.Println(studentGrades)

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Printf("Alice: %+v\n", alice)
	}

	key := "Dan"
	if _, ok := studentGrades[key]; ok {
		fmt.Printf("%s: %+v\n", key, studentGrades[key])
	}
	delete(studentGrades, "Alice")

	fmt.Printf("%+v\n", studentGrades)

	//configs := make(map[string]int)
	//configs := map[string]int{}
	var configs map[string]interface{}
	fmt.Printf("%+v %T\n", configs, configs)

	if configs == nil {
		fmt.Printf("configs is nil\n")
	}
}
