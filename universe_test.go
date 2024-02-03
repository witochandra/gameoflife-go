package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUniverse(t *testing.T) {
	u := NewUniverse([][]int{})
	assert.Equal(t, len(u.CellPositions()), 0)

	u = NewUniverse([][]int{
		{1, 2},
		{1, 2},
		{2, 3},
	})
	assert.Equal(t, len(u.CellPositions()), 2)

	u = NewUniverse([][]int{
		{1, 2},
		{2, 3},
	})
	assert.Equal(t, len(u.CellPositions()), 2)
}
