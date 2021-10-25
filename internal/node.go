package internal

import (
	"math"
	"sync"
)

type node struct {
	lock         sync.Mutex
	children     []node
	centerOfMass location
	totalMass    float64
	width        float64
	location     location
	body         *body
}

func (n *node) addBody(b *body) {
	if n.isLeaf() {
		if n.hasBody() {
			n.convertToInternal()
			n.addBodyToChild(b)
		} else {
			n.body = b
		}
	} else {
		n.addBodyToChild(b)
	}
}

func (n *node) addBodyToChild(b *body) {
	i := n.locationToChildrenIndex(b.position)
	n.children[i].addBody(b)
}

func (n *node) contains(b *body) bool {
	return b.position.x < n.width && b.position.x > 0 && b.position.y < n.width && b.position.y > 0
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (n *node) hasBody() bool {
	return n.body != nil
}

func (n *node) convertToInternal() {
	childLocations := childLocations(n.location, n.width)
	for i := 0; i < 4; i++ {
		n.children = append(n.children, node{location: childLocations[i], width: n.width / 2})
	}

	pTmp := n.body
	n.body = nil
	n.addBody(pTmp)
}

// [nw, ne, sw, se]
func childLocations(l location, w float64) []location {
	return []location{
		location{x: l.x, y: l.y + w/2},
		location{x: l.x + w/2, y: l.y + w/2},
		location{x: l.x, y: l.y},
		location{x: l.x + w/2, y: l.y},
	}
}

// Internal nodes have an array called children which has four elements
// (since this is a quadtree). This function returns the index of the
// child node that would contain the provided location.
func (n *node) locationToChildrenIndex(l location) int {
	if l.x > n.location.x+n.width/2 {
		if l.y > n.location.y+n.width/2 {
			return 1
		}
		return 3
	}

	if l.y > n.location.y+n.width/2 {
		return 0
	}

	return 2
}

func (n *node) calculateCentersOfMass() {
	if n.isLeaf() && n.hasBody() {
		n.totalMass = n.body.mass
		n.centerOfMass = n.body.position
	}

	if !n.isLeaf() {
		for i := range n.children {
			n.children[i].calculateCentersOfMass()
		}

		x := 0.0
		y := 0.0
		totalMass := 0.0

		for i := range n.children {
			totalMass += n.children[i].totalMass
			x += n.children[i].centerOfMass.x * n.children[i].totalMass
			y += n.children[i].centerOfMass.y * n.children[i].totalMass
		}

		n.totalMass = totalMass
		n.centerOfMass = location{x: x / totalMass, y: y / totalMass}
	}
}

func (n *node) calculateForceOnBody(b *body) {
	if n.isLeaf() {
		if n.hasBody() && n.body != b {
			b.addForce(n)
		}
	} else {
		threshold := n.width / math.Sqrt(math.Pow(n.centerOfMass.x-b.position.x, 2)+math.Pow(n.centerOfMass.y-b.position.y, 2))
		if threshold < theta {
			b.addForce(n)
		} else {
			for i := range n.children {
				n.children[i].calculateForceOnBody(b)
			}
		}
	}
}
