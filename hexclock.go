// A simple hextime clock
// See: https://en.wikipedia.org/wiki/Florence_Mean_Time
package main

import (
	"github.com/chrissexton/hexclock/hextime"
	"fmt"
)

func main() {
	fmt.Println(hextime.Now())
}
