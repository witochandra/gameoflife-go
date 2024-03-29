package gameoflife

import "strings"

const (
	DeadCell  = '_'
	AliveCell = '#'
)

type Points []Point

func (ps Points) Edges() (topLeft, bottomRight Point) {
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

func (ps Points) String(offset int, emptyPoint, cellPoint string) string {
	topLeft, bottomRight := ps.Edges()

	lines := make([]string, 0)
	if len(ps) > 0 {
		pointByStr := make(map[string]int)
		for _, p := range ps {
			pointByStr[p.String()] = 1
		}
		for y := topLeft.Y; y <= bottomRight.Y; y++ {
			row := ""
			for x := topLeft.X; x <= bottomRight.X; x++ {
				pos := Point{X: x, Y: y}
				if pointByStr[pos.String()] == 1 {
					row += cellPoint
				} else {
					row += emptyPoint
				}
			}
			lines = append(lines, row)
		}
	}
	if offset > 0 {
		for i, line := range lines {
			lines[i] = strings.Repeat(emptyPoint, offset) + line + strings.Repeat(emptyPoint, offset)
		}
		length := (bottomRight.X - topLeft.X) + (2 * offset)
		if len(lines) > 0 {
			length += 1
		}
		topDown := strings.Repeat(emptyPoint, length)
		for i := 0; i < offset; i++ {
			lines = append([]string{topDown}, lines...)
			lines = append(lines, topDown)
		}
	}

	return strings.Join(lines, "\n") + "\n"
}

func NewPoints(string, cellPoint string) Points {
	lines := strings.Split(string, "\n")
	points := make(Points, 0)
	for y, line := range lines {
		for x, c := range line {
			if c == rune(cellPoint[0]) {
				points = append(points, Point{X: x, Y: y})
			}
		}
	}
	return points
}
