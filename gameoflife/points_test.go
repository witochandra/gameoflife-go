package gameoflife_test

import (
	"gameoflife/gameoflife"
	"testing"
)

func TestPoints_Edges(t *testing.T) {
	t.Run("Empty points", func(t *testing.T) {
		ps := gameoflife.Points{}
		bottomLeft, topRight := ps.Edges()
		if bottomLeft != (gameoflife.Point{}) {
			t.Errorf("Expected {0, 0}, got %v", bottomLeft)
		}

		if topRight != (gameoflife.Point{}) {
			t.Errorf("Expected {0, 0}, got %v", topRight)
		}
	})

	t.Run("Not empty points", func(t *testing.T) {
		ps := gameoflife.Points{
			{X: 1, Y: 1},
			{X: -2, Y: 2},
			{X: 2, Y: 1},
		}

		bottomLeft, topRight := ps.Edges()
		if bottomLeft != (gameoflife.Point{X: -2, Y: 1}) {
			t.Errorf("Expected {-2, 1}, got %v", bottomLeft)
		}

		if topRight != (gameoflife.Point{X: 2, Y: 2}) {
			t.Errorf("Expected {2, 2}, got %v", topRight)
		}
	})
}
