package rule

const (
	leftParenthesis  = "("
	rightParenthesis = ")"
	and              = "$and"
	or               = "$or"
	atLeast          = "$atLeast"
	atMost           = "$atMost"
)

const (
	leafNode = iota
	andNode
	orNode
	atLeastNode
	atMostNode
	parenthesisNode
)

var nodeTypeMap = map[string]int{
	and:             andNode,
	or:              orNode,
	atLeast:         atLeastNode,
	atMost:          atMostNode,
	leftParenthesis: parenthesisNode,
}

// ASTNode 是抽象语法树的节点
type ASTNode struct {
	Child *ASTNode
	// 列表用链表形式存储
	Next     *ASTNode
	Token    string
	NodeType int
}

// appendChild 给当前节点增加一个子节点
func (node *ASTNode) appendChild(anotherNode *ASTNode) {
	if node.Child == nil {
		node.Child = anotherNode
		return
	}
	current := node.Child
	for current.Next != nil {
		current = current.Next
	}
	current.Next = anotherNode
}

// buildAST 递归构造抽象语法树
func buildAST(tks *Tokens) *ASTNode {
	current := tks.peak()
	if current == leftParenthesis {
		node := new(ASTNode)
		node.NodeType = parenthesisNode
		for tks.shift(); tks.peak() != rightParenthesis; {
			node.appendChild(buildAST(tks))
		}
		// 弹出右括号
		tks.shift()
		return node
	}
	tks.shift()
	nodeType, _ := nodeTypeMap[current]
	return &ASTNode{
		Token:    current,
		NodeType: nodeType,
	}
}
