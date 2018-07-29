package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// RequireOnce node
type RequireOnce struct {
	Meta     meta.Collection
	Position *position.Position
	Expr     node.Node
}

// NewRequireOnce node constructor
func NewRequireOnce(Expression node.Node) *RequireOnce {
	return &RequireOnce{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *RequireOnce) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *RequireOnce) GetPosition() *position.Position {
	return n.Position
}

func (n *RequireOnce) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *RequireOnce) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *RequireOnce) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
