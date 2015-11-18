# 001-GA-WebGL

available [gh-page](http://zombiehippie.github.io/Go-Concepts/Play/001-GA-WebGL/)

Here, we have established the very basis of how we can use Golang to interface with a 3D visuallization library.

[`./driver.js`](./driver.js) shows how we hook into Threejs to make interfacing between Golang extremely simplified for our application, because we really don't want to write anymore graphics code than necessary in Go.

`driver.js` does three main things:
 1. it exposes `driver.addEntity(pos_x, pos_y, pos_z, rot_x, rot_y, rot_z, scale_x, scale_y, scale_z, color)`
 2. it exposes `driver.update()` which removes all previous generation geometry from scene, and adds the new generation geometry.
 3. it calls `driver.iter` on interval, which is each interval we iterate our generations, add entities, and update. 