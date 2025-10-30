package parser

import (
	"calculator/lexer"
	"calculator/util"
)

type ASTNode struct {
	Token               lexer.Token
	Depth               uint8
	Left, Right, Parent *ASTNode
}

func (node *ASTNode) IsHigherPriority(other *ASTNode) bool {
	if other.Token.Type == lexer.UNTYPED {
		return false
	}
	if node.Depth != other.Depth {
		return node.Depth > other.Depth
	}
	return node.Token.Type.Priority() > other.Token.Type.Priority()
}

func (ast *ASTNode) MaxDepth() uint8 {
	if ast == nil {
		return 0
	}
	return 1 + max(ast.Left.MaxDepth(), ast.Right.MaxDepth())
}

// The idea is to append the incoming child
// to the right of the current parent.
// And then check if the parent has a higher priority.
// If it is, then I move the child into the parent
// and the parent becomes the left child of the
// 2 + 3 * 4
//
// current node: +
//
//		  +
//		 /
//	  2
//
// current node: 3
//
//		 +
//	  / \
//	 2   3
//
// current node: *
//
//		 +
//	  / \
//	 2   *
//	    /
//	   3
//
// This is the expected result
//
//		 +
//	  / \
//	 2   *
//	    / \
//	   3   4
func Parse(tokens []lexer.Token) *ASTNode {
	// UNTYPED token has the lowest priority as possible
	// such that the root remain at the top
	// avoiding swap a node with the root.
	root := &ASTNode{Token: lexer.Token{Type: lexer.UNTYPED}}

	parent := root
	var depth uint8 = 0

	iter := util.NewIterator(tokens)
	for iter.HasNext() {
		token := iter.Peek()
		if token.Type == lexer.LPAREN {
			depth++
			continue
		}
		if token.Type == lexer.RPAREN {
			depth--
			continue
		}

		node := &ASTNode{Token: token, Depth: depth}
		dummy := parent
		for dummy.Parent != nil {
			if node.IsHigherPriority(dummy) {
				break
			}
			dummy = dummy.Parent

		}
		node.Parent = dummy
		node.Left = dummy.Right
		dummy.Right = node
		if node.Left != nil {
			node.Left.Parent = node
		}

		parent = node

		iter.Next()
	}

	if depth != 0 {
		panic("Mismatch parenthesis")
	}

	return root.Right
}

func (ast *ASTNode) Evaluate() int {
	if ast.Token.Type == lexer.PLUS {
		return ast.Left.Evaluate() + ast.Right.Evaluate()
	}
	if ast.Token.Type == lexer.MINUS {
		return ast.Left.Evaluate() - ast.Right.Evaluate()
	}
	if ast.Token.Type == lexer.STAR {
		return ast.Left.Evaluate() * ast.Right.Evaluate()
	}
	if ast.Token.Type == lexer.SLASH {
		return ast.Left.Evaluate() / ast.Right.Evaluate()
	}
	return ast.Token.Value.(int)
}
