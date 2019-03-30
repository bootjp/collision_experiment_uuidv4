package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/satori/go.uuid"
)

func main() {
	max := runtime.NumCPU()
	s := &sync.Map{}

	for i := 0; i < max; i++ {
		go run(s)
	}

	c := make(chan byte)
	wait := <-c
	fmt.Println(wait)
}

func run(s *sync.Map) {
	for i := 0; i < 10; i++ {

		u1 := uuid.Must(uuid.NewV4())
		st := u1.String()

		_, ok := s.Load(st)

		if ok {
			println("collision uuid!")
			os.Exit(1)
		}
		s.Store(st, st)

	}
}
