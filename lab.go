package main

import (
	"errors"
	"strconv"
	"strings"
)

func Initials(names []string) []string {
	initials := []string{}
	for _, n := range names {
		initials = append(initials, string(n[0]))
	}
	return initials
}

func SnakeCase(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

func Swap(first, last string) (string, string) {
	return last, first
}

func DoubleMultiplesOf(numbers []int, n int) []int {
	doubles := []int{}
	for _, v := range numbers {
		if v%n == 0 {
			doubles = append(doubles, 2*v)
		} else {
			doubles = append(doubles, v)
		}
	}
	return doubles
}

func Filter(strings []string, f func(string) bool) []string {
	filtered := []string{}
	for _, s := range strings {
		if f(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func Inc(counter *int) {
	*counter++
}

func Fibonacci() func() int {
	one, two := 1, 1
	return func() (next int) {
		next = one
		one = two
		two = next + one
		return
	}
}

type Sayer interface {
	say() string
}

type Muter interface {
	mute()
}

type Tapir struct {
	muted bool
}

func (t Tapir) say() string {
	if t.muted {
		return "nada"
	}
	return "oiiiiiiii"
}

func (t *Tapir) mute() {
	t.muted = true
}

type Cow struct {
	muted bool
}

func (c Cow) say() string {
	if c.muted {
		return "nada"
	}
	return "Mooo"
}

func (c *Cow) mute() {
	c.muted = true
}

func FailIfNegative(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Negative values are not allowed: " + strconv.Itoa(n))
	}
	return n, nil
}
