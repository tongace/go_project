package main

import (
	"fmt"
	"time"

	"github.com/tongace/go_project/fibonaci"
)

func main() {
	start1 := time.Now()
	v := fibonaci.CalFibonaciWithRecursive(10)
	elapsed1 := time.Since(start1)
	fmt.Printf("fibonaci of recursive value of 10 is %d \n", v)
	fmt.Printf("time eclibaced1 is %s \n", elapsed1)

	start2 := time.Now()
	f := fibonaci.CalFibonaciWithClosure()
	for ii := 0; ii < 10; ii++ {
		fmt.Println(f())
	}
	elapsed2 := time.Since(start2)
	fmt.Printf("fibonaci of closure value of 10 is %d \n", v)
	fmt.Printf("time eclibaced2 is %s \n", elapsed2)
}
