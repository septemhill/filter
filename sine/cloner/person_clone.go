package cloner

import "github.com/septemhill/misc/sine/gene"

func PersonClone(g gene.PersonGene) gene.PersonGene {
	ng := g.Clone()
	ng.SerialId++
	return ng
}
