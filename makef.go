package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"n,omitempty"`
	Age  int    `json:"a,omitempty"`
}
type E struct {
	Person
	EmployeeId string
}

func main() {
	p := E{Person: Person{Name: "Wanwan", Age: 10},
		EmployeeId: "12345E"}
	fmt.Println(p.Name)
	fmt.Println(p.EmployeeId)

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	var p2 Person
	jsonStr := `{"n":"Marley","a":25}`
	err = json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Unmarshaled struct %+v\n", p2)

}
