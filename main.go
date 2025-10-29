package main

import (
	"alice/lexer"
	"alice/parser"
)

func main() {
	t2 := lexer.Token{Value: 2, Type: lexer.NUMBER}
	t3 := lexer.Token{Value: 3, Type: lexer.NUMBER}
	t4 := lexer.Token{Value: 4, Type: lexer.NUMBER}
	plus := lexer.Token{Value: nil, Type: lexer.PLUS}
	star := lexer.Token{Value: nil, Type: lexer.STAR}
	root := parser.Parse([]lexer.Token{t2, plus, t3, star, t4})

	root.Print()
}
