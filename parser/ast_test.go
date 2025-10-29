package parser

import (
	"alice/lexer"
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.NUMBER, Value: 2},
		{Type: lexer.PLUS, Value: nil},
		{Type: lexer.NUMBER, Value: 3},
	}

	root := Parse(tokens) // la tua funzione parser

	// Verifica la struttura dell'albero
	if root.Token.Type != lexer.PLUS {
		t.Errorf("Expected root to be PLUS, got %v", root.Token.Type)
	}

	if root.Left.Token.Value != 2 {
		t.Errorf("Expected left child to be 2, got %v", root.Left.Token.Value)
	}

	if root.Right.Token.Value != 3 {
		t.Errorf("Expected right child to be 3, got %v", root.Right.Token.Value)
	}
}

func TestPrecedence(t *testing.T) {
	// 2 + 3 * 4
	tokens := []lexer.Token{
		{Type: lexer.NUMBER, Value: 2},
		{Type: lexer.PLUS, Value: nil},
		{Type: lexer.NUMBER, Value: 3},
		{Type: lexer.STAR, Value: nil},
		{Type: lexer.NUMBER, Value: 4},
	}

	root := Parse(tokens)

	// Root dovrebbe essere +
	if root.Token.Type != lexer.PLUS {
		t.Errorf("Expected root to be PLUS, got %v", root.Token.Type)
	}

	// Il figlio destro dovrebbe essere *
	if root.Right.Token.Type != lexer.STAR {
		t.Errorf("Expected right child to be STAR, got %v", root.Right.Token.Type)
	}
}

func TestParentheses(t *testing.T) {
	// (2 + 3) * 4
	tokens := []lexer.Token{
		{Type: lexer.LPAREN, Value: nil},
		{Type: lexer.NUMBER, Value: 2},
		{Type: lexer.PLUS, Value: nil},
		{Type: lexer.NUMBER, Value: 3},
		{Type: lexer.RPAREN, Value: nil},
		{Type: lexer.STAR, Value: nil},
		{Type: lexer.NUMBER, Value: 4},
	}

	root := Parse(tokens)

	// Root dovrebbe essere *
	if root.Token.Type != lexer.STAR {
		t.Errorf("Expected root to be STAR, got %v", root.Token.Type)
	}

	// Il figlio sinistro dovrebbe essere +
	if root.Left.Token.Type != lexer.PLUS {
		t.Errorf("Expected left child to be PLUS, got %v", root.Left.Token.Type)
	}
}
