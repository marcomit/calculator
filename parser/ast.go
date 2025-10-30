package parser

import (
	"calculator/lexer"
	"fmt"
	"math"
)

type ASTNodeType int

const (
	AST_NODE_BINARY  ASTNodeType = iota
	AST_NODE_LITERAL             = iota
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

func (ast *ASTNode) Print() {
	queue := []*ASTNode{}
	queue = append(queue, ast)
	depth := math.Pow(2, float64(ast.MaxDepth()))
	level := 0
	for len(queue) > 0 {
		s := len(queue)
		spaces := float64(int(depth)-s) / float64(s+1)
		str := ""
		for range int(spaces) {
			str += " "
		}
		for range s {
			node := queue[0]
			queue = queue[1:]
			fmt.Print(str)

			if node != nil {
				fmt.Print(node.Token.ToString())
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			} else {
				print(" ")
			}
		}
		level++
		fmt.Println()
	}
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

	for _, token := range tokens {
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
