package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	x := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(x)
}
