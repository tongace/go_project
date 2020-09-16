package main

import (
	"fmt"

	"github.com/tongace/go_project/change"
)

func main() {
	bc := change.Change(1788.80)
	for _, v := range bc {
		fmt.Printf("result of change of Bank Note %.2f = %d \n", v.BankNote, v.Amount)
	}
}
