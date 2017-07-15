package main

import (
	"fmt"

	"sklrsn.github.com/Robot/store"
)

func main() {
	msg := store.SayHello()
	fmt.Println(msg)
}
