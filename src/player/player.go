package player

type Player struct {
	Name string
	X, Y float32
	W, H int16
}

func (p Player) GetPos() (float32, float32) {
	return p.X, p.Y
}

func (p *Player) SetPos(x float32, y float32) {
	p.X = x
	p.Y = y
}