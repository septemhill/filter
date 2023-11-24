package cloner

import (
	"sine/gene"
	"sine/mut"
)

type ClonerFactory[T gene.Gene[T]] interface {
	Create() Cloner[T]
}

type clonerFactory[T gene.Gene[T]] struct {
	clnFn CloneFunc[T]
	muts  []mut.Mutation[T]
}

func (clnf *clonerFactory[T]) Create() Cloner[T] {
	return NewDefaultCloner(clnf.clnFn, clnf.muts...)
}

func NewClonerFactory[T gene.Gene[T]](clnFn CloneFunc[T], muts ...mut.Mutation[T]) *clonerFactory[T] {
	return &clonerFactory[T]{
		clnFn: clnFn,
		muts:  muts,
	}
}
