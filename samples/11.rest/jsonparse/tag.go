package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var jsonString = []byte(`
[
    {"title": "The Art of Community", "author": "Jono Bacon"},
    {"title": "Mithril", "author": "Yoshiki Shibukawa", "price": 1600}
]`)

func main() {
	var books []Book
	err := json.Unmarshal(jsonString, &books)
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}
}
