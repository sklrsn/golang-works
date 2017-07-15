package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"sklrsn.github.com/Robot/store"
)

func main() {
	connection := store.Connection{}
	connection.Initialize("postgres", "love", "postgres")

	product := store.Product{Name: "Bat", Price: 27}
	fmt.Println(store.Persist(&product, &connection))
}
