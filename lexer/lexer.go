package lexer

import (
	"calculator/util"
	"fmt"
	"strconv"
	"unicode"
)

type TokenType uint8

const (
	UNTYPED TokenType = iota
	LPAREN            = iota
	RPAREN            = iota
	NUMBER            = iota
	PLUS              = iota
	MINUS             = iota
	STAR              = iota
	SLASH             = iota
)

type Token struct {
	Type  TokenType
	Value any
}

func (t TokenType) Priority() int {
	switch t {
	case PLUS, MINUS:
		return 1
	case STAR, SLASH:
		return 2
	default:
		return 255
	}
}

func (token *Token) ToString() string {
	switch token.Type {
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case STAR:
		return "*"
	case SLASH:
		return "/"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case UNTYPED:
		return "null"
	}
	return fmt.Sprintf("%v", token.Value)
}

func Tokenize(content string) ([]Token, error) {
	tokens := []Token{}

	iter := util.NewIterator([]rune(content))

	for iter.HasNext() {
		if unicode.IsDigit(iter.Peek()) {
			acc := ""
			for iter.HasNext() && unicode.IsDigit(iter.Peek()) {
				acc += string(iter.Peek())
				iter.Next()
			}
			n, _ := strconv.Atoi(acc)
			tokens = append(tokens, Token{Type: NUMBER, Value: n})
		}
		switch iter.Peek() {
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Value: "+"})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Value: "-"})
		case '*':
			tokens = append(tokens, Token{Type: STAR, Value: "*"})
		case '/':
			tokens = append(tokens, Token{Type: SLASH, Value: "/"})
		}
		iter.Next()
	}

	return tokens, nil
}
