// carriage-return is a terminal process indicator with a check mark
package main

import (
	"fmt"
	"time"
)

func main() {
	s := "............................"
	l := len(s)

	for i := 0; i < l; i++ {
		fmt.Printf("\r[%v] %v/%v ", s, i, l)
		s = s[:i] + "#" + s[i+1:]
		time.Sleep(500 * time.Millisecond)
	}

	// \u2713 prints a (check mark) to stdout
	fmt.Printf("\r[%v] %v/%v \u2713", s, l, l)
	fmt.Println()
}
