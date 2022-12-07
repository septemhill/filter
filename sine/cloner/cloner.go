package cloner

import (
	"github.com/septemhill/misc/sine/gene"
	"github.com/septemhill/misc/sine/mut"
)

type Cloner[T gene.Gene[T]] interface {
	Clone(T) T
}

type CloneFunc[T gene.Gene[T]] func(T) T

type defaultCloner[T gene.Gene[T]] struct {
	clnFn CloneFunc[T]
	muts  []mut.Mutation[T]
}

func (dcln *defaultCloner[T]) Clone(g T) T {
	cg := dcln.clnFn(g)

	for _, m := range dcln.muts {
		cg = *m.Mutate(&cg)
	}

	return cg
}

func NewDefaultCloner[T gene.Gene[T]](clnFn CloneFunc[T], muts ...mut.Mutation[T]) *defaultCloner[T] {
	return &defaultCloner[T]{
		clnFn: clnFn,
		muts:  muts,
	}
}
