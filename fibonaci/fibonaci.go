package fibonaci

//CalFibonaciWithRecursive is function to calculate fibonaci value of n
func CalFibonaciWithRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return CalFibonaciWithRecursive(n-1) + CalFibonaciWithRecursive(n-2)
}

//CalFibonaciWithClosure is function to calculate fibonaci value of n
func CalFibonaciWithClosure() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}
