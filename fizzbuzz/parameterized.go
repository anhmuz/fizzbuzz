package fizzbuzz

import (
	"io"
	"strconv"
)

type parameterized struct {
	params []Parameter
}

type Parameter struct {
	Divider int
	Word    string
}

func NewParameterized(params []Parameter) FizzBuzz {
	return parameterized{params: params}
}

func (fb parameterized) Print(writer io.StringWriter, start int, end int) {
	for i := start; i <= end; i++ {
		printed := false
		for _, param := range fb.params {
			if i%param.Divider == 0 {
				writer.WriteString(param.Word)
				printed = true
			}
		}

		if !printed {
			writer.WriteString(strconv.Itoa(i))
		}

		writer.WriteString("\n")
	}
}
