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
	var names = []string{"Kant", "Popper", "Aristotle", "Socrates", "Plato"}
	var actual = Initials(names)
	var expected = []string{"K", "P", "A", "S", "P"}
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
	var numbers = []int{1, 3, 7, 9, 14, 22, 28}
	var actual = DoubleMultiplesOf(numbers, 7)
	var expected = []int{1, 3, 14, 9, 28, 22, 56}
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
