package lexer

import "fmt"

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

type Token struct {
	Type  TokenType
	Value any
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
