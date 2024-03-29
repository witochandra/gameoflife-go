package gameoflife

type Points []Point

func (ps Points) Edges() (bottomLeft, topRight Point) {
	set := false
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for _, p := range ps {
		if !set {
			minX, minY, maxX, maxY = p.X, p.Y, p.X, p.Y
			set = true
			continue
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return Point{minX, minY}, Point{maxX, maxY}
}
