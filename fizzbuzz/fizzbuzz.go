package fizzbuzz

import "io"

type FizzBuzz interface {
	Print(writer io.StringWriter, start int, end int)
}
