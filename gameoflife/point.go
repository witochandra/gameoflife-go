package gameoflife

type Point struct {
	X, Y int
}

func (p Point) Neighbors() []Point {
	return []Point{
		{p.X - 1, p.Y - 1}, // bottom left
		{p.X, p.Y - 1},     // bottom
		{p.X + 1, p.Y - 1}, // bottom right
		{p.X - 1, p.Y},     // left
		{p.X + 1, p.Y},     // right
		{p.X - 1, p.Y + 1}, // top left
		{p.X, p.Y + 1},     // top
		{p.X + 1, p.Y + 1}, // top right
	}
}
