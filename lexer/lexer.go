package lexer

import (
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

type StringIterator struct {
	runes []rune
	index int
}

func (s *StringIterator) peek() rune {
	return s.runes[s.index]
}

func (s *StringIterator) hasPrev() bool {
	return s.index > 0
}
func (s *StringIterator) has() bool {
	return s.index <= len(s.runes)-1
}
func (s *StringIterator) hasNext() bool {
	return s.index < len(s.runes)-1
}
func (s *StringIterator) next() rune {
	if !s.hasNext() {
		panic("Trying to call next at the end of iteration")
	}
	s.index++
	return s.peek()
}
func (s *StringIterator) prev() rune {
	if !s.hasPrev() {
		panic("Trying to call prev at the start of iteration")
	}
	s.index--
	return s.peek()
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

	iter := &StringIterator{runes: []rune(content), index: 0}

	for iter.has() {
		if unicode.IsDigit(iter.peek()) {
			acc := ""
			for unicode.IsDigit(iter.peek()) {
				acc += string(iter.peek())
				if !iter.hasNext() {
					break
				}
				iter.next()
			}
			n, _ := strconv.Atoi(acc)
			tokens = append(tokens, Token{Type: NUMBER, Value: n})
			if !iter.hasNext() {
				break
			}
			iter.next()
			continue
		}
		switch iter.peek() {
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
		if !iter.hasNext() {
			break
		}
		iter.next()
	}

	return tokens, nil
}
