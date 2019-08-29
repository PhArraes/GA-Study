package bob

import ga "../ga"

//GABob Genetic Algorithm of bobs problem
type GABob struct {
	genomaSize         int
	_totalFitnessScore float64
	BestFitnessScore   float64
	_genomes           []ga.Genome
	CurrentGeneration  int
	BobsMap            BobsMap
}

func New() GABob {
	return GABob{};
}

//TotalFitnessScore get TotalFitnessScore
func (ga GABob) TotalFitnessScore() float64 {
	return ga._totalFitnessScore
}

//Genomes get Genomes
func (ga GABob) Genomes() []ga.Genome {
	return ga._genomes
}

//UpdateFitness Update Fitness
func (ga GABob) UpdateFitness() {
	count := len(ga._genomes)
	for index := 0; index < count; index++ {
		genome := ga._genomes[index]
		steps := translateGenome(genome)
		genome.Fitness = ga.BobsMap.TestRoute(steps)
		ga.SetBestFitnessScore(genome)
	}
}

//SetBestFitnessScore Set Genomes
func (ga GABob) SetBestFitnessScore(genome ga.Genome) {
	if genome.Fitness > ga.BestFitnessScore {
		ga.BestFitnessScore = genome.Fitness
	}
}

//SetGenomes Set Genomes
func (ga GABob) SetGenomes(genomes []ga.Genome) {
	ga._genomes = genomes
}

//IncrementGeneration Increment Generation
func (ga GABob) IncrementGeneration() {
	ga.CurrentGeneration++
}

func translateGenome(genome ga.Genome) []int {
	count := len(genome.Bits)
	steps := make([]int, count/2)
	stepIndex := 0
	for index := 0; index < count; index += 2 {
		steps[stepIndex] = (genome.Bits[index] * 1) + (genome.Bits[index+1] * 2)
		stepIndex++
	}
	return steps
}
