# Golang, a Whirlwind Tour

This repository contains supporting files to the presentation [Golang, a
Whirlwind Tour](). It contains tests that may be completed to get a taste of
Go.

## Installation

```
# Install with brew
brew install go
```

If you don't use Homebrew, follow the [installation instructions](https://golang.org/doc/install).


```
# Clone the repository
git clone git@github.com:andersjanmyr/golang-a-whirlwind-tour.git
```

## General Tests

```
cd golang-a-whirlwind-tour
go test
```

Open `lab.go` and `lab_test.go` in an editor. Follow the instructions and fix
the tests.

## Web Tests

```
cd web
go test
```

Open `handlers.go` and `handlers_test.go` in an editor. Follow the instructions and fix
the tests.


## Concurrency

```
cd ../concurrency
go run philosophers.go
```

Open `philosophers.go` in an editor. Read through the code and play around with
it.


