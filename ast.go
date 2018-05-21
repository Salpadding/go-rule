package rule

import(
	"errors"
)

// astNode 是抽象语法树的节点
type astNode struct {
	child *astNode
	// 列表用链表形式存储
	next     *astNode
	token    *token
}

// appendChild 给当前节点增加一个子节点
func (node *astNode) appendChild(anotherNode *astNode) {
	if node.child == nil {
		node.child = anotherNode
		return
	}
	current := node.child
	for current.next != nil {
		current = current.next
	}
	current.next = anotherNode
}

// buildAST 递归构造抽象语法树
func buildAST(tks *tokens) (*astNode, error) {
	if tks.length() == 0{
		return nil, errors.New("unexpected eof")
	}
	current := tks.peak()
	if current.tokenType == leftParenthesisToken {
		node := new(astNode)
		node.token = current
		for tks.shift(); tks.peak().tokenType != rightParenthesisToken; {
			newNode, err := buildAST(tks)
			if err != nil{
				return nil ,err
			}
			node.appendChild(newNode)
		}
		// 弹出右括号
		tks.shift()
		return node, nil
	}
	tks.shift()
	return &astNode{
		token:    current,
	}, nil
}
