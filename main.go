package main

import (
	"embed"
	"fmt"
	"strings"
	//u "github.com/jboursiquot/stringutils"
)

const proverbs = `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`

//go:embed proverbs.txt
var f embed.FS

func main() {
	data, _ := f.ReadFile("proverbs.txt")
	print(string(data))

	lines := strings.Split(string(data), "\n")
	for i, l := range lines {
		c := i + 1
		wc := len(strings.Fields(l))
		fmt.Printf("%d. %s (WC: %d)\n", c, l, wc)

		fmt.Println()
		lines := strings.Split(proverbs, "\n")
		for _, l := range lines {
			fmt.Printf("%s\n", l)
			for k, v := range charCount(l) {
				fmt.Printf("'%s'=%d, ", k, v)
			}
			fmt.Print("\n\n")
		}
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}
	return m
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
