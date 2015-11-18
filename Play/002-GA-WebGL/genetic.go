package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math/rand"
)

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
	hwgoal := "Hello, world!"
	helloWorldGA := &HelloWorldGA{[]string{
		"Gekmo+ xosmd!",
		"Gekln, worle\"",
		"Fello, wosld!",
		"Gello, wprld!",
		"Hello, world!",
	}, hwgoal}
	driver.Call("connect", func() {
		iteration += 1
		addEntity1(.5, .5, rand.Float32(), rand.Float32())
		addEntity2(.5, rand.Float32(), .55, 1.0, rand.Float32(), .55, rand.Float32())
		driver.Call("update")

		// Hello World! example
		console.Call("log", "Iteration:", iteration)
		helloWorldGA.Evolve()
		population := helloWorldGA.GetPopulation()
		for _, chromosome := range population {
			console.Call("log", "Chromosome:", chromosome, HWFitness(hwgoal, chromosome))
		}
	})
}
