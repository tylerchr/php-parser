package name

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// NamePart node
type NamePart struct {
	Meta     meta.Collection
	Position *position.Position
	Value    string
}

// NewNamePart node constructor
func NewNamePart(Value string) *NamePart {
	return &NamePart{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *NamePart) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *NamePart) GetPosition() *position.Position {
	return n.Position
}

func (n *NamePart) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *NamePart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *NamePart) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
