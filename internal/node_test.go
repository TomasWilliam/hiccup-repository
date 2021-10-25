package internal

import (
	"testing"
)

// TODO: Write better tests
func TestNodesSimple(t *testing.T) {
	root := node{location: location{x: 0, y: 0}, width: 10.0}

	b1 := body{position: location{x: 1, y: 1}}
	b2 := body{position: location{x: -1, y: 1}}

	root.addBody(b1)
	root.addBody(b2)

	if root.children[0].body != b2 && root.children[1].body != b1 {
		t.Error()
	}
}

func TestNodesComplex(t *testing.T) {
	root := node{location: location{x: 0, y: 0}, width: 10.0}

	b1 := body{position: location{x: 1, y: 1}}
	b2 := body{position: location{x: 4.5, y: 4.5}}

	root.addBody(b1)
	root.addBody(b2)

	if root.children[1].children[2].body != b1 && root.children[1].children[1].body != b2 {
		t.Error()
	}
}
