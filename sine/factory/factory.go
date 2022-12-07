package factory

import (
	"github.com/septemhill/misc/sine/assembler"
	"github.com/septemhill/misc/sine/gene"
)

type GeneFactory[T gene.Gene[T]] interface {
	CreateAssembleFactory() assembler.AssemblerFactory[T]
	CreateCloneFactory()
}
