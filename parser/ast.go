package parser

import (
	"alice/lexer"
	"fmt"
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

func (ast *ASTNode) Print() {
	queue := []*ASTNode{}
	queue = append(queue, ast)
	// depth := 0
	// func dfs(curr *ASTNode, d int) {
	// 	if curr == nil {
	// 		depth = max(depth, d)
	// 		return
	// 	}
	// 	dfs(curr.Left, d+1)
	// 	dfs(curr.Right, d+1)
	// }
	for len(queue) > 0 {
		s := len(queue)
		for range s {
			node := queue[0]
			queue = queue[1:]
			fmt.Print(node.Token.ToString())
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
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
	root := &ASTNode{}
	root.Left = &ASTNode{}
	var depth uint8 = 0
	parent := root.Left
	for _, token := range tokens {
		fmt.Printf("parsing %v\n", token.ToString())
		if token.Type == lexer.LPAREN {
			depth++
			continue
		}
		if token.Type == lexer.RPAREN {
			depth--
			continue
		}
		node := &ASTNode{
			Token: token,
			Depth: depth,
		}
		dummy := parent
		for dummy.Parent != nil && !node.IsHigherPriority(dummy.Parent) {
			fmt.Printf("curr %v\n", dummy.Token.ToString())
			dummy = dummy.Parent
		}
		fmt.Printf("parent %v\n", dummy.Token.ToString())
		node.Left = dummy.Right
		node.Parent = dummy
		dummy.Right = node
		parent = node
		fmt.Printf("added %v\n", node.Token.ToString())
		if node.Left != nil {
			fmt.Printf("left %v\n", node.Left.Token.ToString())
		}
		if node.Left != nil {
			fmt.Printf("left %v\n", node.Left.Token.ToString())
		}
	}

	return root.Left
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
