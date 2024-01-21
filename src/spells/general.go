package spells

type Destination struct {
	X float32
	Y float32
}

type Spell interface {
	New(Destination, Destination)
	Move()
	Render()
}

var Projectiles = make([]Spell, 100)