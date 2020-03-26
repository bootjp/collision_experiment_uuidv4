package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/satori/go.uuid"
	"os"
	"runtime"
	"sync/atomic"
)


func main() {
	var ops uint64
	max := runtime.NumCPU() * 2
	store := mapset.NewSet()

	for i := 0; i < max; i++ {
		go run(store, &ops)
	}

	c := make(chan byte)
	wait := <-c
	fmt.Println(wait)
}

func run(s mapset.Set, cnt *uint64) {
	for {
		u1 := uuid.NewV4()
		if s.Contains(u1) {
			fmt.Println("collision uuid!")
			fmt.Printf("genereted %d uuid\n", cnt)
			os.Exit(1)
		}
		atomic.AddUint64(cnt, 1)
		fmt.Printf("%d\n", *cnt)
		s.Add(u1)
	}
}
