package assembler

import (
	"sine/gene"
	"sine/mut"
)

type AssemblerFactory[T gene.Gene[T]] interface {
	Create() Assembler[T]
}

type assemblerFactory[T gene.Gene[T]] struct {
	asmFn AssembleFunc[T]
	muts  []mut.Mutation[T]
}

func (asmf *assemblerFactory[T]) Create() Assembler[T] {
	return NewDefaultAssembler(asmf.asmFn, asmf.muts...)
}

func NewAssemblerFactory[T gene.Gene[T]](asmFn AssembleFunc[T], muts ...mut.Mutation[T]) *assemblerFactory[T] {
	return &assemblerFactory[T]{
		asmFn: asmFn,
		muts:  muts,
	}
}
