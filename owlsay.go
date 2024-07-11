package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// owl taken from https://ascii.co.uk/art/owl (thank you!)

func main() {
	args := os.Args[1:]
	input := ""
	if len(args) == 0 {
		input = print_stdin()
	} else {
		input = strings.Join(args, " ")
	}
	wrapped := wrap(input)
	fmt.Println(wrapped)
	owl := `
         \
          \  ^...^          .
            / o,o \    .    \
            |):::(|     \   /-
==============w=w=======(--^-._
                         \_
`
	fmt.Println(owl)
}

func wrap(input string) [][]string {
	// stripped := strings.ReplaceAll(line, "\n", " ")
	lines := strings.Split(input, "\n")
	lines_of_words := make([][]string, len(lines))
	for i, line := range lines {
		lines_of_words[i] = strings.Split(line, " ") 
	}
	return lines_of_words
}

func print_stdin() string {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	str := string(stdin)
	str = strings.TrimSuffix(str, "\n")
	return str
}
