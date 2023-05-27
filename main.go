package main

import (
	"fizzbuzz/fizzbuzz"
	"os"
)

const chunkSize = 6

func main() {
	// fb := fizzbuzz.NewSimpleFizzBuzz()
	// fb.Print(os.Stdout, 1, 20)

	params := []fizzbuzz.Parameter{
		{Divider: 3, Word: "Fizz"},
		{Divider: 5, Word: "Buzz"},
	}
	fb := fizzbuzz.NewParameterized(params)
	fb = fizzbuzz.WithParallel(fb, chunkSize)
	fb = fizzbuzz.WithBuffer(fb)

	fb.Print(os.Stdout, 1, 20)
}
