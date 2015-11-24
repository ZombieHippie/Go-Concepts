package main

import (
	"math/rand"
	"sort"
)

type HWChromosome string

type HWCitizen struct {
	value   HWChromosome
	fitness int
}
type ByFitness []HWCitizen

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].fitness < a[j].fitness }

type HelloWorldGA struct {
	population []HWCitizen
	goal       HWChromosome
}

func MakeHelloWorldGA(chroms []HWChromosome, goal HWChromosome) *HelloWorldGA {
	h := &HelloWorldGA{
		make([]HWCitizen, len(chroms)),
		goal,
	}
	for index, chrom := range chroms {
		h.population[index] = HWCitizen{
			chrom,
			0,
		}
		h.HWFitness(&h.population[index])
	}
	return h
}

func (h *HelloWorldGA) HWFitness(test *HWCitizen) {
	// assume goal and test are same length for now
	total_cost := 0
	for i := 0; i < len(h.goal); i++ {
		dif := int(h.goal[i]) - int(test.value[i])
		total_cost += dif * dif
	}
	test.fitness = total_cost
}

func (h *HelloWorldGA) Crossover(p1, p2 HWCitizen) (n1, n2 HWCitizen) {
	pivot := rand.Intn(len(string(h.goal)))
	n1.value = p1.value[0:pivot] + p2.value[pivot:]
	n2.value = p2.value[0:pivot] + p1.value[pivot:]
	return
}

func (h *HelloWorldGA) Mutate(p1 *HWCitizen) {
	geneIndex := rand.Intn(len(string(p1.value)))
	b := []byte(p1.value)
	b[geneIndex] += byte(rand.Intn(13) - 6)
	p1.value = HWChromosome(b)
}

func (h *HelloWorldGA) GetPopulation() []HWCitizen {
	return h.population
}

func (h HelloWorldGA) calcMedianFitness() (median_fitness int) {
	medianIndex := len(h.population)/2 - 1
	median_fitness = h.population[medianIndex].fitness
	// if even, then average two centers
	if len(h.population)%2 == 0 {
		median_fitness += h.population[medianIndex-1].fitness
		median_fitness /= 2
	}
	return
}
func (h HelloWorldGA) calcAvgFitness() (avg_fitness int) {
	for _, citizen := range h.population {
		avg_fitness += citizen.fitness
	}
	avg_fitness /= len(h.population)
	return
}

func (h *HelloWorldGA) killUnfit(maxfitness int) {
	// sort by fitness, so the first members are first to crossover
	sort.Sort(ByFitness(h.population))
	// kill all who don't have above average fitness
	// find index at which below maxfitness starts
	var index int
	var citizen HWCitizen
	half := len(h.population) / 2
	for index, citizen = range h.population {
		if citizen.fitness > maxfitness || index == half {
			break
		}
	}
	h.population = h.population[:index]
}
func (h *HelloWorldGA) killCount(deaths int) {
	// sort by fitness, so the first members live
	sort.Sort(ByFitness(h.population))
	lifeIndex := len(h.population) - deaths
	h.population = h.population[:lifeIndex]
}

func (h *HelloWorldGA) mate(births int) {
	i := 0
	mutationCount := 0
	livingTotal := len(h.population)
	Debug("test", births, livingTotal)
	for ; i < births-1; i += 2 {
		parent1 := h.population[i%livingTotal]
		parent2 := h.population[(i+1)%livingTotal]
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
		h.population = append(h.population, newborn1, newborn2)
	}
	mutationRatio := float32(mutationCount) / float32(births)
	Debug("Mutations:", mutationCount, "/", births, "=", float32(int(mutationRatio*1e3))*1e-1, "%")
	// if not all spots have been filled, clone the most fit to the end
	if i < births {
		h.population = append(h.population, h.population[0])
	}
}
func (h *HelloWorldGA) Evolve() {
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
	h.mate(newSpots)
}
