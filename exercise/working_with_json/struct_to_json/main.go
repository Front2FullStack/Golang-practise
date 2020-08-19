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
	fmt.Println("Hello World")
	author := Author{Name: "Mr John Doe", Age: 25, Developer: true}
	book := Book{Title: "Learning JSON in golang", Author: author}
	byteArray, err := json.Marshal(book) // json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(byteArray))
}
