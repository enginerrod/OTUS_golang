package main

import (
	"fmt"

	"golang.org/x/example"
)

func main() {
	x := example.Reverse("Hello, OTUS!")
	fmt.Println(x)
}
