package gameoflife_test

import (
	"gameoflife/gameoflife"
	"strings"
	"testing"
)

func TestPoints_Edges(t *testing.T) {
	t.Run("Empty points", func(t *testing.T) {
		ps := gameoflife.Points{}
		topLeft, bottomRight := ps.Edges()
		if topLeft != (gameoflife.Point{}) {
			t.Errorf("Expected {0, 0}, got %v", topLeft)
		}

		if bottomRight != (gameoflife.Point{}) {
			t.Errorf("Expected {0, 0}, got %v", bottomRight)
		}
	})

	t.Run("Not empty points", func(t *testing.T) {
		ps := gameoflife.Points{
			{X: 1, Y: 1},
			{X: -2, Y: 2},
			{X: 2, Y: 1},
		}

		topLeft, bottomRight := ps.Edges()
		if topLeft != (gameoflife.Point{X: -2, Y: 1}) {
			t.Errorf("Expected {-2, 1}, got %v", topLeft)
		}

		if bottomRight != (gameoflife.Point{X: 2, Y: 2}) {
			t.Errorf("Expected {2, 2}, got %v", bottomRight)
		}
	})
}

func TestPoints_String(t *testing.T) {
	testCases := []struct {
		name     string
		points   gameoflife.Points
		offset   int
		expected string
	}{
		{
			name:     "Empty points",
			points:   gameoflife.Points{},
			offset:   0,
			expected: "\n",
		},
		{
			name:     "Empty points with offset",
			points:   gameoflife.Points{},
			offset:   2,
			expected: "____\n____\n____\n____\n",
		},
		{
			name: "Single point",
			points: gameoflife.Points{
				{X: 1, Y: 1},
			},
			offset:   0,
			expected: "#\n",
		},
		{
			name: "Single point with offset",
			points: gameoflife.Points{
				{X: 1, Y: 1},
			},
			offset: 2,
			expected: strings.Join([]string{
				"_____",
				"_____",
				"__#__",
				"_____",
				"_____",
			}, "\n") + "\n",
		},
		{
			name: "Multiple points with offset",
			points: gameoflife.Points{
				{X: -1, Y: -1},
				{X: 1, Y: 1},
			},
			offset: 2,
			expected: strings.Join([]string{
				"_______",
				"_______",
				"__#____",
				"_______",
				"____#__",
				"_______",
				"_______",
			}, "\n") + "\n",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.points.String(tc.offset, "_", "#")
			if tc.expected != actual {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
