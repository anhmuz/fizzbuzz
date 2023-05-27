package fizzbuzz

import (
	"io"
	"strings"
	"sync"
)

type parallel struct {
	inner     FizzBuzz
	chunkSize int
}

func WithParallel(inner FizzBuzz, chunkSize int) FizzBuzz {
	return parallel{inner: inner, chunkSize: chunkSize}
}

func (fb parallel) Print(writer io.StringWriter, start int, end int) {
	chunksNum := (end - start + 1) / fb.chunkSize
	if (end-start+1)%fb.chunkSize != 0 {
		chunksNum++
	}

	var wg sync.WaitGroup
	builders := make([]strings.Builder, chunksNum)
	for i := 0; i < chunksNum; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			s := start + i*fb.chunkSize
			e := s + fb.chunkSize - 1
			if e > end {
				e = end
			}

			fb.inner.Print(&builders[i], s, e)
		}()
	}

	wg.Wait()

	for _, builder := range builders {
		writer.WriteString(builder.String())
	}
}
