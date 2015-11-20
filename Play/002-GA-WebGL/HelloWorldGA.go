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

func (h *HelloWorldGA) mate() {
	new_population := make([]HWCitizen, 0, len(h.population))
	avg_fitness := 0
	for index, citizen := range h.population {
		Debug("new fitness of", citizen.value, citizen.fitness)
		avg_fitness += citizen.fitness
		h.population[index] = citizen
	}
	avg_fitness /= len(h.population)
	// sort by fitness, so the first members are first to crossover
	sort.Sort(ByFitness(h.population))
	// kill all who don't have above average fitness
	for index, citizen := range h.population {
		if citizen.fitness <= avg_fitness {
			new_population = append(new_population, citizen)
			if index > len(h.population)/2 {
				break
			}
		} else {
			break
		}
	}
	// fill the rest of the new_population with crossovers
	livingTotal := len(new_population)
	newSpots := len(h.population) - livingTotal
	Debug("New Spots", newSpots, livingTotal)
	i := 0
	for ; i < newSpots; i += 2 {
		parent1 := new_population[i%livingTotal]
		parent2 := new_population[(i+1)%livingTotal]
		newborn1, newborn2 := h.Crossover(parent1, parent2)
		// 1 in 5 mutate
		if rand.Intn(5) == 0 {
			Debug("Mutate 1")
			h.Mutate(&newborn1)
		}
		if rand.Intn(5) == 0 {
			Debug("Mutate 2")
			h.Mutate(&newborn2)
		}
		h.HWFitness(&newborn1)
		h.HWFitness(&newborn2)
		new_population = append(new_population, newborn1, newborn2)
	}
	if i+1 < newSpots {
		new_population = append(new_population, new_population[0])
	}
	h.population = new_population
}
func (h *HelloWorldGA) Evolve() {
	h.mate()
}
