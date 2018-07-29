package scalar

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// EncapsedStringPart node
type EncapsedStringPart struct {
	Meta     meta.Collection
	Position *position.Position
	Value    string
}

// NewEncapsedStringPart node constructor
func NewEncapsedStringPart(Value string) *EncapsedStringPart {
	return &EncapsedStringPart{
		Value: Value,
	}
}

// SetPosition sets node position
func (n *EncapsedStringPart) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *EncapsedStringPart) GetPosition() *position.Position {
	return n.Position
}

func (n *EncapsedStringPart) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *EncapsedStringPart) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *EncapsedStringPart) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
