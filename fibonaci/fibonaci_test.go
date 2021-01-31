package fibonaci_test

import (
	"testing"

	"github.com/tongace/go_project/fibonaci"
)

func TestWith0mustbe0(t *testing.T) {
	v := fibonaci.CalFibonaciWithSequence(0)
	if v != 0 {
		t.Errorf("n = 0 return value is %d", v)
	}
}
func TestWith1mustbe1(t *testing.T) {
	v := fibonaci.CalFibonaciWithSequence(1)
	if v != 1 {
		t.Errorf("n = 1 return value is %d", v)
	}
}
func TestWith2mustbe1(t *testing.T) {
	v := fibonaci.CalFibonaciWithSequence(2)
	if v != 1 {
		t.Errorf("n = 2 return value is %d", v)
	}
}
func TestWith3mustbe2(t *testing.T) {
	v := fibonaci.CalFibonaciWithSequence(3)
	if v != 2 {
		t.Errorf("n = 3 return value is %d", v)
	}
}
func TestWith4mustbe3(t *testing.T) {
	v := fibonaci.CalFibonaciWithSequence(4)
	if v != 3 {
		t.Errorf("n = 4 return value is %d", v)
	}
}
func TestWith10mustbe55(t *testing.T) {
	v := fibonaci.CalFibonaciWithSequence(10)
	if v != 55 {
		t.Errorf("n = 10 return value is %d", v)
	}
}
func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonaci.CalFibonaciWithRecursive(10)
	}
}

func BenchmarkFibonacciSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonaci.CalFibonaciWithSequence(10)
	}
}

func BenchmarkFibonacciClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := fibonaci.CalFibonaciWithClosure()
		for jj := 0; jj < 10; jj++ {
			f()
		}
		fibonaci.CalFibonaciWithSequence(10)
	}
}
