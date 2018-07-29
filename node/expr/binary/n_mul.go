package binary

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Mul node
type Mul struct {
	Meta     meta.Collection
	Position *position.Position
	Left     node.Node
	Right    node.Node
}

// NewMul node constructor
func NewMul(Variable node.Node, Expression node.Node) *Mul {
	return &Mul{
		Left:  Variable,
		Right: Expression,
	}
}

// SetPosition sets node position
func (n *Mul) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Mul) GetPosition() *position.Position {
	return n.Position
}

func (n *Mul) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *Mul) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Mul) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}
