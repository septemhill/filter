package mut

import "sine/gene"

type Mutation[T gene.Gene[T]] interface {
	Mutate(*T) *T
}
