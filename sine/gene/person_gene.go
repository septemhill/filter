package gene

type PersonGene struct {
	Color    []byte
	SerialId uint64
}

func (pg PersonGene) Type() GeneType {
	return GeneTypePerson
}

func (pg PersonGene) Clone() PersonGene {
	return pg
}

var _ Gene[PersonGene] = (PersonGene{})
