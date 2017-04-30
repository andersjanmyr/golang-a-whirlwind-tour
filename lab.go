package main

func Initials(names []string) []string {
	return []string{}
}

func SnakeCase(name string) string {
	return ""
}

func Swap(first, last string) (string, string) {
	return "", ""
}

func DoubleMultiplesOf(numbers []int, n int) []int {
	return []int{}
}

func Filter(strings []string, f func(string) bool) []string {
	return []string{}
}

func Inc(counter *int) {
}

func Fibonacci() func() int {
	return func() int {
		return 0
	}
}

type Sayer interface {
	say() string
}

type Muter interface {
	mute()
}

type Tapir struct {
}

func (tapir Tapir) say() string {
	return ""
}

func (tapir Tapir) mute() {
}

// Implement Cow here
