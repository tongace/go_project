package main

import (
	"fmt"

	"github.com/tongace/go_project/fizzbuzz"
)

func main() {
	param := 3
	v := fizzbuzz.FizzBuzz(param)
	fmt.Printf("return value of param %d is %s", param, v)
}
