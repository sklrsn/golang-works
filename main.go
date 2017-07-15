package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"sklrsn.github.com/Robot/store"
)

func main() {
	msg := store.SayHello()
	fmt.Println(msg)

	sum := store.Sum(1, 2)
	fmt.Println(sum)

	connection := store.Connection{}
	connection.Initialize("postgres", "love", "postgres")

	person := store.Person{Name: "Kalai", Age: 27}
	store.Persist(&person, &connection)
}
