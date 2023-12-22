package main

import (
	"fmt"
	"time"

	"os"

	"github.com/DMXMax/concurrency/cmd/quotegen"
	"github.com/DMXMax/concurrency/generator"
	"github.com/rs/zerolog/log"
)

type message struct {
	Body      string
	TimeStamp int64
}

func jsonStreamGnerator(done <-chan any) <-chan any {
	stream := make(chan any)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- message{Body: "Hello", TimeStamp: time.Now().UnixNano()}:
			}
		}
	}()
	return stream
}

func main() {
	file, err := os.Create("log.txt")
	defer file.Close()

	if err != nil {
		log.Fatal().Err(err).Msg("Error creating log file")
	}

	done := make(chan any)
	defer close(done)
	for num := range generator.Take(done, generator.Repeat(done, 1, 2), 10) {
		fmt.Println(num)
	}

	for dat := range generator.Take(done, jsonStreamGnerator(done), 10) {
		fmt.Printf("%v\n", dat)
	}

	for quote := range generator.Take(done, quotegen.Generator(done), 10) {
		if quote, ok := quote.(quotegen.Quote); ok {
			file.WriteString(fmt.Sprintf("%#v\n", quote))
		}
	}
}
