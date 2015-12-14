package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math/rand"
)

func Debug(args ...interface{}) {
	js.Global.Get("document").Get("driver").Call("visualLog", args...)
}
func DebugSpecial(args ...interface{}) {
	js.Global.Get("document").Get("driver").Call("visualLogHighlight", args...)
}

func main() {
	rand.Seed(42)
	document := js.Global.Get("document")
	console := js.Global.Get("console")
	driver := document.Get("driver")

	addEntity3 := func(px, py, pz, rx, ry, rz, sx, sy, sz, color float32) {
		driver.Call("addEntity", px, py, pz, rx, ry, rz, sx, sy, sz, color)
	}
	addEntity2 := func(px, py, pz, rx, ry, rz, color float32) {
		addEntity3(px, py, pz, rx, ry, rz, 0, 0, 0, color)
	}
	addEntity1 := func(px, py, pz, color float32) {
		addEntity3(px, py, pz, 0, 0, 0, 0, 0, 0, color)
	}

	iteration := 0
	hwgoal := HWChromosome("Hello, world!")

	populationSize := 256
	testChromosomes := make([]HWChromosome, 0, populationSize)
	for i := 0; i < populationSize; i++ {
		str := make([]byte, len(hwgoal))
		for j := 0; j < len(hwgoal); j++ {
			str[j] = byte(32 + rand.Intn(64)) // general range of bytes
		}
		testChromosomes = append(testChromosomes, HWChromosome(str))
	}
	helloWorldGA := MakeHelloWorldGA(testChromosomes, hwgoal)

	running := true
	driver.Call("connect", func() {
		if running {
			addEntity1(.5, .5, rand.Float32(), rand.Float32())
			addEntity2(.5, rand.Float32(), .55, 1.0, rand.Float32(), .55, rand.Float32())
			driver.Call("update")

			iteration += 1
			// Hello World! example
			console.Call("log", "Iteration:", iteration)
			Debug("Iteration:", iteration)
			helloWorldGA.Evolve()
			population := helloWorldGA.GetPopulation()
			for _, citizen := range population {
				if citizen.fitness == 0 {
					DebugSpecial("Chromosome:", citizen.value, citizen.fitness)
					running = false
				} else {
					Debug("Chromosome:", citizen.value, citizen.fitness)
				}
			}
			driver.Call("updateVisualLog")
		}
	})
}
