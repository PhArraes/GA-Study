package ga

import "math/rand"

//GeneticAlgorithm behavior of a GA
type GeneticAlgorithm interface {
	TotalFitnessScore() float64
	Genomes() []Genome
	UpdateFitness()
	SetGenomes([]Genome)
	IncrementGeneration()
}

//Epoch behavior of a GA
func Epoch(ga GeneticAlgorithm) {
	ga.UpdateFitness()
	ga.SetGenomes(breed(ga))
	ga.IncrementGeneration()
}

//RouletteWheelSelection select a genome from a collection
func rouletteWheelSelection(genomes []Genome, totalFitnessScore float64) Genome {
	slice := rand.Float64() * totalFitnessScore
	count := len(genomes)
	total := 0.0
	selectedGenome := 0
	for i := 0; i < count; i++ {
		total += genomes[i].Fitness
		if total > slice {
			selectedGenome = i
			break
		}
	}
	return genomes[selectedGenome]
}

func breed(ga GeneticAlgorithm) []Genome {
	numBabies := 0
	babyGenomes := make([]Genome, PopulationSize)
	for numBabies < PopulationSize {
		mum := rouletteWheelSelection(ga.Genomes(), ga.TotalFitnessScore())
		dad := rouletteWheelSelection(ga.Genomes(), ga.TotalFitnessScore())
		baby1, baby2 := mum.Crossover(dad)
		baby1.Mutate()
		baby2.Mutate()
		babyGenomes[numBabies] = baby1
		numBabies++
		babyGenomes[numBabies] = baby2
		numBabies++
	}
	return babyGenomes
}
