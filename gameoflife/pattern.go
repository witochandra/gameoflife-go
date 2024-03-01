package gameoflife

func CreateBlinker(x, y int) []Point {
	return []Point{
		{X: x + 2, Y: y + 1},
		{X: x + 2, Y: y + 2},
		{X: x + 2, Y: y + 3},
	}
}

func CreateToad(x, y int) []Point {
	return []Point{
		{X: x + 2, Y: y + 2}, {X: x + 3, Y: y + 2}, {X: x + 4, Y: y + 2},
		{X: x + 1, Y: y + 3}, {X: x + 2, Y: y + 3}, {X: x + 3, Y: y + 3},
	}
}

func CreateBeacon(x, y int) []Point {
	return []Point{
		{X: x + 1, Y: y + 1}, {X: x + 2, Y: y + 1},
		{X: x + 1, Y: y + 2}, {X: x + 2, Y: y + 2},
		{X: x + 3, Y: y + 3}, {X: x + 4, Y: y + 3},
		{X: x + 3, Y: y + 4}, {X: x + 4, Y: y + 4},
	}
}

func CreatePulsar(x, y int) []Point {
	return []Point{
		{X: x + 4, Y: y + 2}, {X: x + 5, Y: y + 2}, {X: x + 6, Y: y + 2}, {X: x + 10, Y: y + 2}, {X: x + 11, Y: y + 2}, {X: x + 12, Y: y + 2},
		{X: x + 2, Y: y + 4}, {X: x + 7, Y: y + 4}, {X: x + 9, Y: y + 4}, {X: x + 14, Y: y + 4},
		{X: x + 2, Y: y + 5}, {X: x + 7, Y: y + 5}, {X: x + 9, Y: y + 5}, {X: x + 14, Y: y + 5},
		{X: x + 2, Y: y + 6}, {X: x + 7, Y: y + 6}, {X: x + 9, Y: y + 6}, {X: x + 14, Y: y + 6},
		{X: x + 4, Y: y + 7}, {X: x + 5, Y: y + 7}, {X: x + 6, Y: y + 7}, {X: x + 10, Y: y + 7}, {X: x + 11, Y: y + 7}, {X: x + 12, Y: y + 7},
		{X: x + 4, Y: y + 9}, {X: x + 5, Y: y + 9}, {X: x + 6, Y: y + 9}, {X: x + 10, Y: y + 9}, {X: x + 11, Y: y + 9}, {X: x + 12, Y: y + 9},
		{X: x + 2, Y: y + 10}, {X: x + 7, Y: y + 10}, {X: x + 9, Y: y + 10}, {X: x + 14, Y: y + 10},
		{X: x + 2, Y: y + 11}, {X: x + 7, Y: y + 11}, {X: x + 9, Y: y + 11}, {X: x + 14, Y: y + 11},
		{X: x + 2, Y: y + 12}, {X: x + 7, Y: y + 12}, {X: x + 9, Y: y + 12}, {X: x + 14, Y: y + 12},
		{X: x + 4, Y: y + 14}, {X: x + 5, Y: y + 14}, {X: x + 6, Y: y + 14}, {X: x + 10, Y: y + 14}, {X: x + 11, Y: y + 14}, {X: x + 12, Y: y + 14},
	}
}

func CreateGlider(x, y int) []Point {
	return []Point{
		{X: x + 3, Y: y + 1},
		{X: x + 1, Y: y + 2}, {X: x + 3, Y: y + 2},
		{X: x + 2, Y: y + 3}, {X: x + 3, Y: y + 3},
	}
}
