package main

import (
	"bufio"
	"fmt"
	"os"
)

// owl taken from https://ascii.co.uk/art/owl (thank you!)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
	owl := `
           ^...^
          / o,o \
          |):::(|
        ====w=w===
`
	fmt.Println(owl)

}
