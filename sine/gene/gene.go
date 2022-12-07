package gene

type GeneType int

const (
	GeneTypePerson = 0
)

type Gene[T any] interface {
	Type() GeneType
	Clone() T
}
