package main

import (
	"fmt"
)

func exchange(exVal float64) {
	var note1000Count int
	if exVal >= 1000 {
		note1000Count = int(exVal / 1000)
		exVal = exVal - float64(note1000Count*1000)
	}
	fmt.Printf("exchange Bank Note of 1,000 amount %d notes", note1000Count)
}

func main() {
	var changeAmount float64
	fmt.Scan()
	exchange(changeAmount)
}
