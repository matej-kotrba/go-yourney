package player

type Player struct {
	Name string
	x, y float32
}

func (p Player) GetPos() (float32, float32) {
	return p.x, p.y
}

func (p *Player) SetPos(x float32, y float32) {
	p.x = x
	p.y = y
}