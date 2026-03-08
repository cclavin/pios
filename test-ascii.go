package main

import (
	"fmt"
	"strings"
)

// To test changes to this ASCII art without recompiling the entire PIOS framework,
// simply edit the string below and run `go run test-ascii.go` in your terminal!
// Once you are happy with the alignment, copy this string back into `cmd/pios/main.go`
const piosAscii = `
    ____  ____  ____  _____
   / __ \/  _// __ \/ ___/ 
  / /_/ // / / / / /\__ \  
 / ____// / / /_/ /___/ /  
/_/   /___/\____//____/    
`

func main() {
	lines := strings.Split(strings.Trim(piosAscii, "\n"), "\n")
	for i, line := range lines {
		if i < len(lines)/2 {
			fmt.Printf("\033[38;5;208m%s\033[0m\n", line) // Orange
		} else {
			fmt.Printf("\033[32m%s\033[0m\n", line) // Green
		}
	}
	fmt.Println()
}
