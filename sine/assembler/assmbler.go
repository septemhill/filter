package assembler

import (
	"sine/gene"
	"sine/mut"
)

type Assembler[T gene.Gene[T]] interface {
	Assemble(*T, *T) *T
}

type AssembleFunc[T gene.Gene[T]] func(*T, *T) *T

type defaultAssembler[T gene.Gene[T]] struct {
	asmFn AssembleFunc[T]
	muts  []mut.Mutation[T]
}

func (dasm *defaultAssembler[T]) Assemble(g1, g2 *T) *T {
	g := dasm.asmFn(g1, g2)

	for _, m := range dasm.muts {
		g = m.Mutate(g)
	}

	return g
}

func NewDefaultAssembler[T gene.Gene[T]](asmFn AssembleFunc[T], muts ...mut.Mutation[T]) *defaultAssembler[T] {
	return &defaultAssembler[T]{
		asmFn: asmFn,
		muts:  muts,
	}
}
