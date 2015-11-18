package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	document := js.Global.Get("document")
	console := js.Global.Get("console")
	driver := document.Get("driver")

	iteration := 0
	driver.Call("connect", func() {
		iteration += 1
		driver.Call("addEntity", .5, .5, .5)
		driver.Call("addEntity", .5, .55, .55)
		driver.Call("update")
		console.Call("log", "Hello", iteration)
	})
}
