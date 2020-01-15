package main

import (
	"fmt"
)

type Table struct {
	Forks        []Fork
	Philosophers []Philosopher
}

type Philosopher struct {
	Id        int
	leftFork  *Fork
	rightFork *Fork
}

type Fork struct {
	Id   int
	Free bool
}

func (p *Philosopher) think() {
	fmt.Printf("Pholosopher %d is thinking", p.Id)
}

func (p *Philosopher) eat() {
	fmt.Printf("Philosopher %d is eating", p.Id)
}

func (p *Philosopher) takeLeftFork(fork *Fork) *Fork {
	if fork.Free != true {
		return nil
	}
	return fork
}

func (p *Philosopher) takeRightFork(fork *Fork) *Fork {
	if fork.Free != true {
		return nil
	}
	return fork
}

func (p *Philosopher) putFork(fork *Fork) *Fork {
	fork.Free = true
	return fork
}

func main() {
	phs := make([]Philosopher, 5)
	for i, _ := range phs {
		phs[i] = Philosopher{
			Id: i,
		}
	}

	forks := make([]Fork, 5)
	for i, _ := range forks {
		forks[i] = Fork{
			Id:   i,
			Free: true,
		}
	}
}
