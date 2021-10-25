package internal

import (
	"math"
)

type body struct {
	position location
	velocity location
	force    location
	mass     float64
	netForce float64
}

func (b *body) addForce(n *node) {
	thetaIJ := math.Atan2(n.centerOfMass.y-b.position.y, n.centerOfMass.x-b.position.x) + math.Pi

	dx := n.centerOfMass.x - b.position.x
	dy := n.centerOfMass.y - b.position.y
	distance := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	F := (g * b.mass * n.totalMass) / math.Pow(distance, 2)

	b.force.x += F * math.Cos(thetaIJ) // dx / distance
	b.force.y += F * math.Sin(thetaIJ) // dy / distance
}

func (b *body) applyForce() {
	b.velocity.x += b.force.x / b.mass
	b.velocity.y += b.force.y / b.mass

	b.position.x -= b.velocity.x
	b.position.y -= b.velocity.y

	b.force = location{}
}
