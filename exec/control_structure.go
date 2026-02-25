package exec

import (
	"github.com/pgcamus/gonja/v2/nodes"
)

type ControlStructure interface {
	nodes.ControlStructure
	Execute(*Renderer, *nodes.ControlStructureBlock) error
}
