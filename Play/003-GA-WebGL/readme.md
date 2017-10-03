# 002-GA-WebGL

available [gh-page](http://colelawrence.github.io/Go-Concepts/Play/003-GA-WebGL/)

On behalf of our good friend [Luke Anderson](github.com/0la0), we have implemented this version of the Hello World genetic algorithm to use probabilistic distribution map to determine the best parents, where a parent is more likely to be picked for reproduction if they have lower cost. Although this seems to take longer than the other method of killing and filling we had in 002, it is much better equipped to prevent getting stuck at local optima.  

We are using gopherjs to compile [HelloWorldGA.go](./HelloWorldGA.go) and [genetic.go](./genetic.go) to `genetic.go.js`.

from 002-GA-WebGL:

Here, we have continued our progress from 001-GA-WebGL to do more in the way of genetic algorithms and moving towards machine learning (ML).

We are using [Burak Kanber's Machine Learning tutorials](http://burakkanber.com/blog/machine-learning-genetic-algorithms-part-1-javascript/) to play here.

from 001-GA-WebGL:

[`./driver.js`](./driver.js) shows how we hook into Threejs to make interfacing between Golang extremely simplified for our application, because we really don't want to write anymore graphics code than necessary in Go.

`driver.js` does three main things:
 1. it exposes `driver.addEntity(pos_x, pos_y, pos_z, rot_x, rot_y, rot_z, scale_x, scale_y, scale_z, color)`
 2. it exposes `driver.update()` which removes all previous generation geometry from scene, and adds the new generation geometry.
 3. it calls `driver.iter` on interval, which is each interval we iterate our generations, add entities, and update. 
