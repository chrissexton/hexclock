// A simple hextime clock
// See: https://en.wikipedia.org/wiki/Florence_Mean_Time
package main

import (
	"flag"
	"fmt"
	"github.com/chrissexton/hexclock/hextime"
)

var (
	format = flag.String(
		"format",
		hextime.BoardmanNotationSeconds,
		"format of time output",
	)
)

func main() {
	flag.Parse()
	fmt.Println(hextime.Nowf(*format))
}
