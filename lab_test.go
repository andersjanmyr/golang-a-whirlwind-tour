package main

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	first, last := Swap("brave", "tapir")
	expectedFirst, expectedLast := "tapir", "brave"
	if first != expectedFirst {
		t.Errorf("Expected %v, actual %v", expectedFirst, first)
	}
	if last != expectedLast {
		t.Errorf("Expected %v, actual %v", expectedLast, last)
	}
}

func TestInitials(t *testing.T) {
	names := []string{"Kant", "Popper", "Aristotle", "Socrates", "Plato"}
	actual := Initials(names)
	expected := []string{"K", "P", "A", "S", "P"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v\n", expected, actual)
	}
}

func TestSnakeCaseDasherized(t *testing.T) {
	actual := SnakeCase("brave-tapir")
	expected := "brave_tapir"
	if actual != expected {
		t.Errorf("Exected %v, actual %v", expected, actual)
	}
}

func TestDoubleMultiplesOf(t *testing.T) {
	numbers := []int{1, 3, 7, 9, 14, 22, 28}
	actual := DoubleMultiplesOf(numbers, 7)
	expected := []int{1, 3, 14, 9, 28, 22, 56}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v\n", expected, actual)
	}
}

func TestFilter(t *testing.T) {
	var names = []string{"Kant", "Popper", "Aristotle", "Socrates", "Plato"}
	var actual = Filter(names, func(name string) bool {
		// Fix filter here
		return true
	})
	var expected = []string{"Popper", "Plato"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v\n", expected, actual)
	}
}

func TestIncrease(t *testing.T) {
	i := 1
	Inc(&i)
	expected := 2
	if i != expected {
		t.Errorf("Expected %v, actual %v", expected, i)
	}
}

func TestMap(t *testing.T) {
	ascii := map[string]int{"A": 65, "B": 67}
	// Write code here
	if !reflect.DeepEqual(map[string]int{"A": 65, "B": 66}, ascii) {
		t.Errorf("Expected %v, actual %v", map[string]int{"A": 65, "B": 66}, ascii)
	}
	// Write code here
	if !reflect.DeepEqual(map[string]int{"A": 65}, ascii) {
		t.Errorf("Expected %v, actual %v", map[string]int{"A": 65}, ascii)
	}
	// Write code here
	if !reflect.DeepEqual(map[string]int{"A": 65, "C": 67}, ascii) {
		t.Errorf("Expected %v, actual %v", map[string]int{"A": 65, "C": 67}, ascii)
	}
}

func assert(expected, actual interface{}, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	assert(1, f(), t)
	assert(1, f(), t)
	assert(2, f(), t)
	assert(3, f(), t)
	assert(5, f(), t)
	assert(8, f(), t)
	assert(13, f(), t)
	assert(21, f(), t)
}

func TestInterfaces(t *testing.T) {
	tapir := Tapir{}
	assert("oiiiiiiii", tapir.say(), t)
	tapir.mute()
	assert("nada", tapir.say(), t)
}

func speak(sayer Sayer) string {
	return sayer.say()
}

func silence(muter Muter) {
	muter.mute()
}

// func TestInterfaces2(t *testing.T) {
// 	cow := Cow{}
// 	assert("Mooo", cow.say(), t)
// 	tapir.mute()
// 	assert("nada", cow.say(), t)
// }
