package main

import (
	"bufio"
	"calculator/lexer"
	"calculator/parser"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		line = strings.TrimSpace(line)
		if line == "exit" {
			fmt.Println("Bye from marcomit")
			return
		}

		tokens, err := lexer.Tokenize(line)
		if err != nil {
			fmt.Print("Error: ", err)
			continue
		}
		root := parser.Parse(tokens)

		fmt.Println(root.Evaluate())
	}
}
