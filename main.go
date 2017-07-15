package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"sklrsn.github.com/Robot/store"
)

var (
	USERNAME = "postgres"
	PASSWORD = "love"
	DATABASE = "postgres"
)

func main() {
	connection := store.Connection{}
	connection.Initialize(USERNAME, PASSWORD, DATABASE)

	product := store.Product{Name: "Bat", Price: 27}
	fmt.Println(store.Persist(&product, &connection))
}
