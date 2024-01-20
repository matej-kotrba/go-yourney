package spells

type Destination struct {
	X float32
	Y float32
}

type Spell interface {
	spawn()
}