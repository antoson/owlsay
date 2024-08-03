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

func wrap(input string) string {
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.ReplaceAll(input, "\t\t", "\n\n")
	words := strings.Split(input, " ")
	var builder strings.Builder
	var charcount = 0
	for _, word := range words {
		// fmt.Printf("%v: %v\n", i, word)
		length, _ := builder.WriteString(word)
		charcount += length
		if charcount >= 30 {
			// fmt.Println("charcount >= 40")
			builder.WriteString("\n")
			charcount = 0
		} else {
			builder.WriteString(" ")
		}
	}
	return builder.String()
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
