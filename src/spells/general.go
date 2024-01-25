package spells

type Destination struct {
	X float32
	Y float32
}

type Spell interface {
	Move()
	Render()
	ShouldBeDeleted() bool
}

var Projectiles = make([]Spell, 100)