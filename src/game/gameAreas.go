package game

type GameArea struct {
	IsActive bool
}

func (ga *GameArea) SetArea(isActive bool) {
	ga.IsActive = isActive
}
