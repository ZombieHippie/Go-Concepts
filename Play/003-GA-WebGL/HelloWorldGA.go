package main

import (
	"math/rand"
)

type HWChromosome string

type HWIndividual struct {
	value   HWChromosome
	fitness int
}
type ByFitness []HWIndividual

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].fitness < a[j].fitness }

type HelloWorldGA struct {
	population []HWIndividual
	goal       HWChromosome
}

// used for making distribution
type probabilisticDistribution struct {
	totalRange  float64
	ranges      []float64
	individuals []HWIndividual
}

func (p *probabilisticDistribution) addChance(individual HWIndividual, weight float64) {
	p.totalRange += weight
	p.ranges = append(p.ranges, p.totalRange)
	p.individuals = append(p.individuals, individual)
}

func (p *probabilisticDistribution) selectRandom() HWIndividual {
	rangeRand := rand.Float64() * p.totalRange
	var lastIndex int
	var rang float64
	for lastIndex, rang = range p.ranges {
		if rang > rangeRand {
			break
		}
	}
	return p.individuals[lastIndex]
}

func MakeHelloWorldGA(chroms []HWChromosome, goal HWChromosome) *HelloWorldGA {
	h := &HelloWorldGA{
		make([]HWIndividual, len(chroms)),
		goal,
	}
	for index, chrom := range chroms {
		h.population[index] = HWIndividual{
			chrom,
			0,
		}
		h.HWFitness(&h.population[index])
	}
	return h
}

func (h *HelloWorldGA) HWFitness(test *HWIndividual) {
	// assume goal and test are same length for now
	total_cost := 0
	for i := 0; i < len(h.goal); i++ {
		dif := int(h.goal[i]) - int(test.value[i])
		total_cost += dif * dif
	}
	test.fitness = total_cost
}

func (h *HelloWorldGA) Crossover(p1, p2 HWIndividual) (n1, n2 HWIndividual) {
	pivot := rand.Intn(len(string(h.goal)))
	n1.value = p1.value[0:pivot] + p2.value[pivot:]
	n2.value = p2.value[0:pivot] + p1.value[pivot:]
	return
}

func (h *HelloWorldGA) Mutate(p1 *HWIndividual) {
	geneIndex := rand.Intn(len(string(p1.value)))
	b := []byte(p1.value)
	b[geneIndex] += byte(rand.Intn(13) - 6)
	p1.value = HWChromosome(b)
}

func (h *HelloWorldGA) GetPopulation() []HWIndividual {
	return h.population
}

func (h *HelloWorldGA) mate() {
	var parents probabilisticDistribution
	for _, ind := range h.population {
		weight := 100.0 / float64(ind.fitness)
		parents.addChance(ind, weight)
	}

	mutationCount := 0
	popSize := len(h.population)
	newPopulation := make([]HWIndividual, 0, popSize)
	for i := 0; i < popSize; i += 2 {
		parent1 := parents.selectRandom()
		parent2 := parents.selectRandom()
		newborn1, newborn2 := h.Crossover(parent1, parent2)
		// 1 in 5 births mutate
		if rand.Intn(5) == 0 {
			h.Mutate(&newborn1)
			mutationCount += 1
		}
		if rand.Intn(5) == 0 {
			h.Mutate(&newborn2)
			mutationCount += 1
		}
		h.HWFitness(&newborn1)
		h.HWFitness(&newborn2)
		newPopulation = append(newPopulation, newborn1, newborn2)
	}
	mutationRatio := float32(mutationCount) / float32(popSize)
	Debug("Mutations:", mutationCount, "/", popSize, "=", float32(int(mutationRatio*1e3))*1e-1, "%")
	h.population = newPopulation
}
func (h *HelloWorldGA) Evolve() {
	/*
		// use average fitness and median fitness to kill off
		avg_fitness := h.calcAvgFitness()
		median_fitness := h.calcMedianFitness()
		Debug("Average fitness:", avg_fitness)
		Debug("Median fitness:", median_fitness)

		populationSize := len(h.population)

		// kill citizens with below threshold fitness
		// or bottom half; at least half the population should die.
		h.killUnfit(avg_fitness)
		// fill the rest of the population with crossovers
		livingTotal := len(h.population)
		newSpots := populationSize - livingTotal
		Debug("Parents, Born, Total:", livingTotal, newSpots, newSpots+livingTotal)
	*/
	h.mate()
}
