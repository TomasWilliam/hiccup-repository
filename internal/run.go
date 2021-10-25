package internal

import (
	"math/rand"
	"time"
)

var width = 400.0
var n = 300
var ticks = 250
var g = 6.67e-4 // e-11
var theta = 0.0

// Run the simulation
func Run() {
	bodies := generateRandomBodies(n)
	locationData := [][]location{}

	for t := 0; t < ticks; t++ {
		//locationData = append(locationData, extractLocations(bodies))
		locationData = append(locationData, extractLocations(generateSpecificBodies()))
		root := constructQuadtree(bodies)
		root.calculateCentersOfMass()
		calculateAndApplyForces(root, bodies)
	}

	generateGif(locationData)
}

func generateRandomBodies(n int) []body {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomLocation := func() location {
		return location{
			x: r.Float64() * width,
			y: r.Float64() * width,
		}
	}

	points := []body{}

	for i := 0; i < n; i++ {
		points = append(points, body{position: randomLocation(), mass: r.Float64() * 10000})
	}

	return points
}

func generateSpecificBodies() []body {
	return []body{
		body{
			position: location{
				x: 400.0,
				y: 400.0,
			},
			mass: 100,
		},
		body{
			position: location{
				x: 200.0,
				y: 400.0,
			},
			mass: 100,
		},
		body{
			position: location{
				x: 400.0,
				y: 500.0,
			},
			mass: 100,
		},
	}
}

func extractLocations(bodies []body) []location {
	locations := []location{}
	for i := range bodies {
		locations = append(locations, bodies[i].position)
	}

	return locations
}

func constructQuadtree(bodies []body) *node {
	root := node{location: location{x: 0, y: 0}, width: width}
	for i := range bodies {
		if root.contains(&bodies[i]) {
			root.addBody(&bodies[i])
		}
	}

	return &root
}

func calculateAndApplyForces(root *node, bodies []body) {
	for i := range bodies {
		if root.contains(&bodies[i]) {
			root.calculateForceOnBody(&bodies[i])
		}
	}

	for i := range bodies {
		bodies[i].applyForce()
	}
}
