package main

import (
	"bufio"
	"fmt"
	"glo/perror"
	"os"
)

var hadError bool = true

type Token struct {
	name string
}

func main() {

	numArgs := len(os.Args[1:])
	if numArgs > 1 {
		fmt.Print("Usage: glox [script]")
		os.Exit(64)
	} else if numArgs == 1 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}

}

func runFile(path string) {
	byte, err := os.ReadFile(path)
	perror.Check(err)

	run(string(byte))

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")
		line, _, err := reader.ReadLine()
		perror.Check(err)
		run(string(line))
	}
}

func run(source string) {
	fmt.Print(source)
	//scanner := Scanner(source)
	//tokens := []Token{scanner.scanTokens()}
	//for i := 0; i<len(tokens); i++ {
	//	fmt.Println(tokens[i].name)
	//}
}
