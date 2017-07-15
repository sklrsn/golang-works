package main

import (
	"fmt"

	store "sklrsn.github.com/Robot/store"
)

func main() {
	msg := store.sayHello()
	fmt.Println(msg)
}
