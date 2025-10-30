package parser

import (
	"calculator/lexer"
	"testing"
)

func TestNestedParentheses(t *testing.T) {
	t1 := lexer.Token{Value: 1, Type: lexer.NUMBER}
	t2 := lexer.Token{Value: 2, Type: lexer.NUMBER}
	t3 := lexer.Token{Value: 3, Type: lexer.NUMBER}
	plus := lexer.Token{Type: lexer.PLUS}
	star := lexer.Token{Type: lexer.STAR}
	lparen := lexer.Token{Type: lexer.LPAREN}
	rparen := lexer.Token{Type: lexer.RPAREN}

	// ((1 + 2) * 3)
	tokens := []lexer.Token{
		lparen, lparen, t1, plus, t2, rparen, star, t3, rparen,
	}

	root := Parse(tokens)
	got := root.Evaluate()
	want := 9 // (1 + 2) * 3

	if got != want {
		t.Errorf("Evaluate() = %v, want %v", got, want)
	}

	if depth := root.MaxDepth(); depth < 2 {
		t.Errorf("MaxDepth() = %v, want >= 2", depth)
	}
}

func TestSingleNumber(t *testing.T) {
	t5 := lexer.Token{Value: 5, Type: lexer.NUMBER}
	root := Parse([]lexer.Token{t5})

	if got := root.Evaluate(); got != 5 {
		t.Errorf("Evaluate() = %v, want 5", got)
	}

	if depth := root.MaxDepth(); depth != 1 {
		t.Errorf("MaxDepth() = %v, want 1", depth)
	}
}

func TestSimpleExpression(t *testing.T) {
	t2 := lexer.Token{Value: 2, Type: lexer.NUMBER}
	t3 := lexer.Token{Value: 3, Type: lexer.NUMBER}
	t4 := lexer.Token{Value: 4, Type: lexer.NUMBER}
	plus := lexer.Token{Type: lexer.PLUS}
	star := lexer.Token{Type: lexer.STAR}
	lparen := lexer.Token{Type: lexer.LPAREN}
	rparen := lexer.Token{Type: lexer.RPAREN}

	root := Parse([]lexer.Token{lparen, t2, plus, t3, rparen, star, t4})

	got := root.Evaluate()
	want := 20 // (2 + 3) * 4

	if got != want {
		t.Errorf("Evaluate() = %v, want %v", got, want)
	}

	maxDepth := root.MaxDepth()
	if maxDepth <= 0 {
		t.Errorf("MaxDepth() = %v, want positive", maxDepth)
	}
}
