package main

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
)


func main() {
	var (
		bufferSize  = 2000
		bufferLimit = 1900
	)
	buf := make([]byte, 0, bufferSize)

	file, err := os.Create("/tmp/uuid_single")
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
}
