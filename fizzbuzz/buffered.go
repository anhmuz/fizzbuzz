package fizzbuzz

import (
	"io"
	"strings"
)

type buffered struct {
	inner FizzBuzz
}

func WithBuffer(inner FizzBuzz) FizzBuzz {
	return buffered{inner: inner}
}

func (fb buffered) Print(writer io.StringWriter, start int, end int) {
	var builder strings.Builder
	fb.inner.Print(&builder, start, end)
	writer.WriteString(builder.String())
}
