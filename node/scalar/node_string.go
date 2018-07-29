package scalar

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// String node
type String struct {
	Meta     meta.Collection
	Position *position.Position
	Value    string
}

// NewString node constructor
func NewString(Value string) *String {
	return &String{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *String) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *String) GetPosition() *position.Position {
	return n.Position
}

func (n *String) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *String) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *String) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
