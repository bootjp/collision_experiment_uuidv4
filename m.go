package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	const thread = 16
	for i := uint64(3); i < uint64(2*thread+3); i += 2 {
		go prime(i, thread)
	}
	time.Sleep(time.Minute)
}

func prime(n uint64, thread int) {
	for i := n; i < math.MaxUint64; i += uint64(2 * thread) {
		isPrime := true
		for j := uint64(3); j*j <= i; j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			//fmt.Println(i)
		}
		if i > 1000000000 {
			fmt.Println(i, n)
		}
	}
}
