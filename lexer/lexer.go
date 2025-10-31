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
	types := make(map[rune]TokenType)
	types['+'] = PLUS
	types['-'] = MINUS
	types['*'] = STAR
	types['/'] = SLASH
	tokens := []Token{}
	iter := util.NewIterator([]rune(content))

	for iter.HasNext() {
		curr := iter.Next()

		if unicode.IsDigit(curr) {
			acc := string(curr)
			for iter.HasNext() && unicode.IsDigit(iter.Peek()) {
				acc += string(iter.Next())
			}
			n, _ := strconv.Atoi(acc)
			tokens = append(tokens, Token{Type: NUMBER, Value: n})
			continue
		}
		switch curr {
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
		case '+', '-', '*', '/':
			tokens = append(tokens, Token{Type: types[curr], Value: string(curr)})
		case ' ':
			continue
		default:
			return []Token{}, fmt.Errorf("Invalid token %v", string(curr))
		}
	}
	return tokens, nil
}
