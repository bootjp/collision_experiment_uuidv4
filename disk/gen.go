package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	uuid "github.com/satori/go.uuid"
)

const (
	bufferSize  = 2000
	bufferLimit = 1900
)

func main() {
	max := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	for i := 0; i < max; i++ {
		wg.Add(1)
		go run(wg, i)
	}

	wg.Wait()
	fmt.Println("DONE")
}

func run(wg *sync.WaitGroup, my int) {
	buf := make([]byte, 0, bufferSize)

	file, err := os.Create(fmt.Sprintf("/tmp/uuid_%d", my))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	for {
		uid := uuid.NewV4()
		buf = append(append(buf, uid.String()...), '\n')
		if len(buf) > bufferLimit {
			_, err := file.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
			buf = make([]byte, 0, bufferSize)

		}
	}
	//wg.Done()
}
