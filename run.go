package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"os"
	"runtime"

	"github.com/satori/go.uuid"
)


func main() {
	max := runtime.NumCPU() * 2
	store := mapset.NewSet()

	for i := 0; i < max; i++ {
		go run(store)
	}

	c := make(chan byte)
	wait := <-c
	fmt.Println(wait)
}

func run(s mapset.Set) {
	for {

		u1 := uuid.NewV4()

		if s.Contains(u1) {
			println("collision uuid!")
			os.Exit(1)
		}
		fmt.Println(u1.String())
		s.Add(u1)
	}
}
