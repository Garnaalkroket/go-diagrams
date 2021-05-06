package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type cdContainer struct {
	path  string
	attrs []attr.Attribute
}

var Cd = &cdContainer{path: "assets/apps/cd"}

func (c *cdContainer) Tekton(opts ...attr.Attribute) *node.Node {
	return node.New("tekton", attr.AssetImage("assets/apps/cd/tekton.png"), attr.Shape(attr.None))
}

func (c *cdContainer) Spinnaker(opts ...attr.Attribute) *node.Node {
	return node.New("spinnaker", attr.AssetImage("assets/apps/cd/spinnaker.png"), attr.Shape(attr.None))
}

func (c *cdContainer) TektonCli(opts ...attr.Attribute) *node.Node {
	return node.New("tekton-cli", attr.AssetImage("assets/apps/cd/tekton-cli.png"), attr.Shape(attr.None))
}
