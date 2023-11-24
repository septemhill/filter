package factory

import (
	"sine/assembler"
	"sine/gene"
)

type GeneFactory[T gene.Gene[T]] interface {
	CreateAssembleFactory() assembler.AssemblerFactory[T]
	CreateCloneFactory()
}
