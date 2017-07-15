package main

import (
	"fmt"

	"sklrsn.github.com/Robot/store"
)

func main() {
	msg := store.SayHello()
	fmt.Println(msg)

	sum := store.Sum(1, 2)
	fmt.Println(sum)
}
