package gameoflife_test

import (
	"gameoflife/gameoflife"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	t.Run("Any live cell with fewer than two live neighbors dies, as if by underpopulation.", func(t *testing.T) {
		initialState := []gameoflife.Point{
			{0, 0},
		}

		assert.Empty(t, gameoflife.Next(initialState))

		initialState = []gameoflife.Point{
			{0, 0}, {0, 1},
		}

		assert.Empty(t, gameoflife.Next(initialState))
	})

	t.Run("Any live cell with more than three live neighbors dies, as if by overpopulation.", func(t *testing.T) {
		/*
			x x      x
			xx   ->  x
			 xx      xxx
		*/
		initialState := []gameoflife.Point{
			{0, 0}, {2, 0},
			{0, -1}, {1, -1},
			{1, -2}, {2, -2},
		}
		assert.ElementsMatch(t,
			[]gameoflife.Point{
				{0, 0},
				{0, -1},
				{0, -2}, {1, -2}, {2, -2},
			},
			gameoflife.Next(initialState),
		)
	})
}
