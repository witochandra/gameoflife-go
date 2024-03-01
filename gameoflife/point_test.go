package gameoflife_test

import (
	"gameoflife/gameoflife"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_Neighbors(t *testing.T) {
	neighbors := gameoflife.Point{0, 0}.Neighbors()
	assert.ElementsMatch(t, []gameoflife.Point{
		{-1, -1}, // bottom left
		{0, -1},  // bottom
		{1, -1},  // bottom right
		{-1, 0},  // left
		{1, 0},   // right
		{-1, 1},  // top left
		{0, 1},   // top
		{1, 1},   // top right
	}, neighbors)
}
