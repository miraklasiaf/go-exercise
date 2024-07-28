// Modify the echo program to print the index and value of each its arguments, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args {
		s += sep + arg
		sep = " "
		fmt.Println(i, " ", arg)
	}
}
