package name

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// FullyQualified node
type FullyQualified struct {
	Meta     meta.Collection
	Position *position.Position
	Parts    []node.Node
}

// NewFullyQualified node constructor
func NewFullyQualified(Parts []node.Node) *FullyQualified {
	return &FullyQualified{
		Parts: Parts,
	}
}

// SetPosition sets node position
func (n *FullyQualified) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *FullyQualified) GetPosition() *position.Position {
	return n.Position
}

func (n *FullyQualified) GetMeta() *meta.Collection {
	return &n.Meta
}

// Attributes returns node attributes as map
func (n *FullyQualified) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *FullyQualified) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		v.EnterChildList("Parts", n)
		for _, nn := range n.Parts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Parts", n)
	}

	v.LeaveNode(n)
}

// GetParts returns the name parts
func (n *FullyQualified) GetParts() []node.Node {
	return n.Parts
}
