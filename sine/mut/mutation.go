package mut

import "github.com/septemhill/misc/sine/gene"

type Mutation[T gene.Gene[T]] interface {
	Mutate(*T) *T
}
