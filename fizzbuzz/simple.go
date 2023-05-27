package fizzbuzz

import (
	"io"
	"strconv"
)

type simpleFizzBuzz struct {
}

func NewSimpleFizzBuzz() FizzBuzz {
	return simpleFizzBuzz{}
}

func (fb simpleFizzBuzz) Print(writer io.StringWriter, start int, end int) {
	for i := start; i <= end; i++ {
		if i%3 != 0 && i%5 != 0 {
			writer.WriteString(strconv.Itoa(i))
		}

		if i%3 == 0 {
			writer.WriteString("Fizz")
		}

		if i%5 == 0 {
			writer.WriteString("Buzz")
		}

		writer.WriteString("\n")
	}
}
