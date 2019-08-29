package ga

import (
	"math/rand"

	utils "../utils"
)

//Evolve class
type Evolve interface {
	Mutate()
	Crossover(dad Genome) (Genome, Genome)
}

//Genome class
type Genome struct {
	Bits    []int
	Fitness float64
}

//New Constructs Genome with a count
func New(count int) Genome {
	bits := make([]int, count)
	for i := 0; i < count; i++ {
		bits[i] = rand.Intn(2)
	}
	return Genome{Fitness: 0, Bits: bits}
}

//Mutate Mutate the genome
func (genome Genome) Mutate() {
	count := len(genome.Bits)
	for i := 0; i < count; i++ {
		if rand.Float64() < MutationRate {
			genome.Bits[i] = (genome.Bits[i] + 1) % 2
		}
	}
}

//Crossover crossover 2 genomes and bread 2 babies
func (genome Genome) Crossover(dad Genome) (Genome, Genome) {
	if rand.Float64() > CrossoverRate || utils.IntArrayEquals(genome.Bits, dad.Bits) {
		return genome, dad
	}
	cp := rand.Intn(GenomeSize - 1)
	baby1, baby2 := New(GenomeSize), New(GenomeSize)
	for i := 0; i < cp; i++ {
		baby1.Bits[i] = genome.Bits[i]
		baby2.Bits[i] = dad.Bits[i]
	}
	for i := cp; i < GenomeSize; i++ {
		baby1.Bits[i] = genome.Bits[i]
		baby2.Bits[i] = dad.Bits[i]
	}
	return baby1, baby2
}
