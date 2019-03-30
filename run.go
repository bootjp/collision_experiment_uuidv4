package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/satori/go.uuid"
)

func main() {
	max := runtime.NumCPU() * 2
	s := &sync.Map{}

	for i := 0; i < max; i++ {
		go run(s)
	}

	c := make(chan byte)
	wait := <-c
	fmt.Println(wait)
}

func run(s *sync.Map) {
	for {

		u1 := uuid.Must(uuid.NewV4())
		_, ok := s.Load(u1)

		if ok {
			println("collision uuid!")
			os.Exit(1)
		}
		fmt.Println(u1.String())
		s.Store(u1, nil)
	}
}
