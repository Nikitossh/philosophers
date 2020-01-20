package main

import (
	"fmt"
	"sync"
)


type SafeCounter struct {
	count	int
	mux	sync.Mutex
}

func NewCounter() SafeCounter {
	s := SafeCounter{
		count: 0,
	}
	return s
}

type Philosopher struct {
	Id        int
}

func NewPhilosopher(sc *SafeCounter) Philosopher {
	sc.IncCounter()
	return Philosopher{
		Id: sc.count,
	}
}

func (sc *SafeCounter) IncCounter() {
        sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.count++
        fmt.Printf("Count: %d", sc.count)
}

func (p *Philosopher) String() {
       fmt.Println(p.Id) 
}

func main() {
        sc := NewCounter()
        phils := make([]Philosopher, 10)
	for i := 0; i < 10; i++ {
	        phils[i] = NewPhilosopher(&sc)
	}
	fmt.Printf("Last Count: %d", sc.count)
}
