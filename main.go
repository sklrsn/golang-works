package main

import (
	"fmt"

	store "sklrsn.github.com/Robot/store"
)

func main() {
	msg := store.SayHello()
	fmt.Println(msg)
}
