package core

type Collider struct {
	position Vector2
	width    int
	height   int
	enabled  bool
}

func IsCollided(self Collider, other Collider) bool {
	// TODO: assumes position.Y is at the top - might be incorrect
	selfLeft := float64(self.position.X)
	selfRight := self.position.X + float64(self.width)
	selfTop := float64(self.position.Y)
	selfBottom := self.position.Y - float64(self.height)

	otherLeft := float64(other.position.X)
	otherRight := other.position.X + float64(other.width)
	otherTop := float64(other.position.Y)
	otherBottom := other.position.Y - float64(other.height)

	if selfLeft < otherRight &&
		selfRight > otherLeft &&
		selfTop > otherBottom &&
		selfBottom < otherTop {
		return true
	}

	return false
}
