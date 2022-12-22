package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Println("Hey there", name)
}

func main() {
	greet("Jörg")
}
