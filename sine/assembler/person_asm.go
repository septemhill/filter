package assembler

import "github.com/septemhill/misc/sine/gene"

func PersonAssemble(g1, g2 *gene.PersonGene) *gene.PersonGene {
	b := make([]byte, len(g1.Color))

	for i := 0; i < len(g1.Color); i++ {
		if i%2 == 1 {
			b[i] = g1.Color[i]
		} else {
			b[i] = g2.Color[i]
		}
	}

	return &gene.PersonGene{
		Color: b,
	}
}
