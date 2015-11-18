package main

type HelloWorldGA struct {
	population []string
	goal       string
}

func HWFitness(goal, test string) int {
	// assume goal and test are same length for now
	total_cost := 0
	for i := 0; i < len(goal); i++ {
		dif := int(goal[i]) - int(test[i])
		total_cost += dif * dif
	}
	return total_cost
}

func (h *HelloWorldGA) GetPopulation() []string {
	return h.population
}

func (h *HelloWorldGA) Evolve() {

}
