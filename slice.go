package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/satori/go.uuid"
)

type result struct {
	mutex sync.RWMutex
	Item  []uuid.UUID
}

func (r *result) Add(u uuid.UUID) {
	if r == nil {
		panic("")
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.Item = append(r.Item, u)
}

func (r *result) Has(u uuid.UUID) bool {
	if r == nil {
		panic("")
	}
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, v := range r.Item {
		a, err := v.MarshalBinary()
		if err != nil {
			panic(err)
		}
		b, err := u.MarshalBinary()
		if err != nil {
			panic(err)
		}

		if bytes.Equal(a, b) {
			return true
		}
	}

	return false

}

func main() {
	res := &result{}
	max := runtime.NumCPU()
	for i := 0; i < max; i++ {
		go runA(res)
	}
	c := make(chan byte)
	wait := <-c
	fmt.Println(wait)

}

func runA(res *result) {
	for {

		u1 := uuid.NewV4()

		if res.Has(u1) {
			println("collision uuid! %s", u1.String())
			os.Exit(1)
		}
		fmt.Println(u1.String())

		res.Add(u1)
	}
}
