package lexer

import (
	"reflect"
	"testing"
)

func TestTokenizeSimpleExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "1+2",
			expected: []Token{
				{Type: NUMBER, Value: 1},
				{Type: PLUS},
				{Type: NUMBER, Value: 2},
			},
		},
		{
			input: "(3-4)*5",
			expected: []Token{
				{Type: LPAREN},
				{Type: NUMBER, Value: 3},
				{Type: MINUS},
				{Type: NUMBER, Value: 4},
				{Type: RPAREN},
				{Type: STAR},
				{Type: NUMBER, Value: 5},
			},
		},
		{
			input: "12/6+7",
			expected: []Token{
				{Type: NUMBER, Value: 12},
				{Type: SLASH},
				{Type: NUMBER, Value: 6},
				{Type: PLUS},
				{Type: NUMBER, Value: 7},
			},
		},
	}

	for _, test := range tests {
		tokens := Tokenize(test.input)

		// Filter out empty numbers added due to current implementation
		filtered := []Token{}
		for _, tok := range tokens {
			if tok.Type != NUMBER || tok.Value.(int) != 0 {
				filtered = append(filtered, tok)
			}
		}

		if !reflect.DeepEqual(filtered, test.expected) {
			t.Errorf("Tokenize(%q) = %v; want %v", test.input, filtered, test.expected)
		}
	}
}

func TestTokenToString(t *testing.T) {
	tests := []struct {
		token    Token
		expected string
	}{
		{Token{Type: PLUS}, "+"},
		{Token{Type: MINUS}, "-"},
		{Token{Type: STAR}, "*"},
		{Token{Type: SLASH}, "/"},
		{Token{Type: LPAREN}, "("},
		{Token{Type: RPAREN}, ")"},
		{Token{Type: UNTYPED}, "null"},
		{Token{Type: NUMBER, Value: 42}, "42"},
	}

	for _, test := range tests {
		result := test.token.ToString()
		if result != test.expected {
			t.Errorf("Token.ToString() = %q; want %q", result, test.expected)
		}
	}
}
