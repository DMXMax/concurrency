package main

import (
	"fmt"

	"github.com/DMXMax/concurrency/generator"
)

func main() {
	done := make(chan any)
	defer close(done)
	for num := range generator.Take(done, generator.Repeat(done, 1, 2), 10) {
		fmt.Println(num)
	}
}
