package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	jsonString := `{"title": "Learning JSON in golang", "author": { "name": "Mr John Doe", "age": 25, "is_developer": true}}`

	var data Book // Also, we can use map[string]interface{} but not recommended
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", data)

	// output if map[string]interface{} used
	// map[author:map[age:25 is_developer:true name:Mr John Doe] title:Learning JSON in golang]
}
