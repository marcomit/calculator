package main

import (
	"calculator/lexer"
	"calculator/parser"
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	println("Press 'exit' to quit the calculator ")
	rl, err := readline.New(">>> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "exit" {
			fmt.Println("Bye from marcomit")
			return
		}

		tokens, err := lexer.Tokenize(line)
		if err != nil {
			fmt.Print("Error: ", err, "\n")
			continue
		}
		root, err := parser.Parse(tokens)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		root.Print()
		println()

		fmt.Println(root.Evaluate())
	}
	fmt.Println("Bye from marcomit")
}
