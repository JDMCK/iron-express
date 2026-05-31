package gui

type Rectangle struct {
	x, y          int
	width, height int
}

func (r *Rectangle) InBounds(x, y int) bool {
	return x >= r.x && x <= r.x+r.width && y >= r.y && y <= r.y+r.height
}
