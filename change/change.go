package change

// Bc :Banknote and Coin
type Bc struct {
	BankNote float64
	Amount   int
}

var bcUnit []Bc

func initBcUnit() {
	bcUnit = []Bc{
		Bc{1000, 0},
		Bc{500, 0},
		Bc{100, 0},
		Bc{50, 0},
		Bc{20, 0},
		Bc{10, 0},
		Bc{5, 0},
		Bc{2, 0},
		Bc{1, 0},
		Bc{0.5, 0},
		Bc{0.25, 0},
	}
}

// Change is function to Change Money of Thai baht
func Change(exVal float64) []Bc {
	initBcUnit()
	var retBC []Bc
	for _, v := range bcUnit {
		if exVal >= v.BankNote {
			v.Amount = int(exVal / v.BankNote)
			exVal = exVal - float64(float64(v.Amount)*v.BankNote)
		}
		retBC = append(retBC, v)
	}
	return retBC
}
