package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	leftFork  chan bool
	rightFork chan bool
	name      string
}

var names = []string{"Kant", "Popper", "Aristotle", "Socrates", "Plato"}
var philosopherCount = len(names)

func main() {
	var forks []chan bool
	var philosophers []Philosopher
	for i := 0; i < philosopherCount; i++ {
		forks = append(forks, makeFork(i))
	}
	for i := 0; i < philosopherCount; i++ {
		philosophers = append(philosophers, makePhilosopher(names[i], i, forks))
	}
	fmt.Println("Press enter to exit")
	time.Sleep(time.Second)
	for i := 0; i < philosopherCount; i++ {
		go philosophers[i].dine()
	}
	var input string
	fmt.Scanln(&input)
}

func makeFork(i int) chan bool {
	fork := make(chan bool, 1)
	fork <- true
	return fork
}

func makePhilosopher(name string, i int, forks []chan bool) Philosopher {
	size := len(forks)
	return Philosopher{forks[i%size], forks[(i+1)%size], name}
}

func (phil *Philosopher) dine() {
	for {
		phil.think()
		phil.takeForks()
		phil.eat()
		phil.putDownForks()
	}
}

func (phil *Philosopher) think() {
	fmt.Printf("%v is thinking.\n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (phil *Philosopher) takeForks() {
	<-phil.leftFork
	fmt.Printf("%v picked up left fork.\n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9))) //Speeds up deadlock
	select {
	case <-phil.rightFork:
		fmt.Printf("%v picked up right fork.\n", phil.name)
		return
	case <-time.After(time.Second):
		fmt.Printf("%v put down left fork.\n", phil.name)
		phil.leftFork <- true
	}
}

func (phil *Philosopher) eat() {
	fmt.Printf("%v is eating.\n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (phil *Philosopher) putDownForks() {
	fmt.Printf("%v put down both forks.\n", phil.name)
	phil.rightFork <- true
	phil.leftFork <- true
}
