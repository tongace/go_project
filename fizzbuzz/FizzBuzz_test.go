package fizzbuzz

import (
	"testing"
)

func TestInput0ShouldbeDisplay0(t *testing.T) {
	v := FizzBuzz(0)
	if v != "0" {
		t.Error("fizzbuzz of 0 should be '0' but have", v)
	}
}
func TestInput1ShouldbeDisplay1(t *testing.T) {
	v := FizzBuzz(1)
	if v != "1" {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}
func TestInput2ShouldbeDisplay2(t *testing.T) {
	v := FizzBuzz(2)
	if v != "2" {
		t.Error("fizzbuzz of 2 should be '2' but have", v)
	}
}
func TestInput3ShouldbeDisplayFizz(t *testing.T) {
	v := FizzBuzz(3)
	if v != "Fizz" {
		t.Error("fizzbuzz of 3 should be 'Fizz' but have", v)
	}
}
func TestInput4ShouldbeDisplay4(t *testing.T) {
	v := FizzBuzz(4)
	if v != "4" {
		t.Error("fizzbuzz of 4 should be '4' but have", v)
	}
}
func TestInput5ShouldbeDisplayBuzz(t *testing.T) {
	v := FizzBuzz(5)
	if v != "Buzz" {
		t.Error("fizzbuzz of 5 should be 'Buzz' but have", v)
	}
}

func TestInput6ShouldbeDisplayFizz(t *testing.T) {
	v := FizzBuzz(6)
	if v != "Fizz" {
		t.Error("fizzbuzz of 6 should be 'Fizz' but have", v)
	}
}
func TestInput7ShouldbeDisplay7(t *testing.T) {
	v := FizzBuzz(7)
	if v != "7" {
		t.Error("fizzbuzz of 7 should be '7' but have", v)
	}
}
func TestInput9ShouldbeDisplayFizz(t *testing.T) {
	v := FizzBuzz(9)
	if v != "Fizz" {
		t.Error("fizzbuzz of 9 should be 'Fizz' but have", v)
	}
}
func TestInput10ShouldbeDisplayFizz(t *testing.T) {
	v := FizzBuzz(10)
	if v != "Buzz" {
		t.Error("fizzbuzz of 10 should be 'Buzz' but have", v)
	}
}
func TestInput15ShouldbeDisplayFizzBuzz(t *testing.T) {
	v := FizzBuzz(15)
	if v != "FizzBuzz" {
		t.Error("fizzbuzz of 15 should be 'FizzBuzz' but have", v)
	}
}
