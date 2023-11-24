package cloner

import "sine/gene"

func PersonClone(g gene.PersonGene) gene.PersonGene {
	ng := g.Clone()
	ng.SerialId++
	return ng
}
