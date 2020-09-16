package change

import (
	"fmt"
)

type bc struct {
	BankNote float64
	Amount   int
}

var bcUnit []bc

func initBcUnit() {
	bcUnit = []bc{
		bc{1000, 0},
		bc{500, 0},
		bc{100, 0},
		bc{50, 0},
		bc{20, 0},
		bc{10, 0},
		bc{5, 0},
		bc{2, 0},
		bc{1, 0},
		bc{0.5, 0},
		bc{0.25, 0},
	}
}

// Change is function to Change Money of Thai baht
func Change(exVal float64) {
	initBcUnit()
	for _, v := range bcUnit {
		if exVal >= v.BankNote {
			v.Amount = int(exVal / v.BankNote)
			exVal = exVal - float64(float64(v.Amount)*v.BankNote)
		}
		fmt.Printf("result of change of Bank Note %d = %d \n", int(v.BankNote), v.Amount)
	}
}
