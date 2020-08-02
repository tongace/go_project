package main

import (
	"fmt"

	"github.com/tongace/go_project/a"
)

func SayB() {
	fmt.Println("Hey !!! Say B")
}

func SayAll() {
	a.SayA()
	SayB()
}

func main() {
	SayAll()
}
