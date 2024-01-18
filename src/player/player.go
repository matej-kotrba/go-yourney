package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	ga "github.com/matej-kotrba/go-testing/src/game"
)

// playerImage = rl.

type Player struct {
	Name  string
	Image *rl.Image
	Texture rl.Texture2D
	X, Y  float32
	W, H  int16
	AreaX int8
	AreaY int8
}

func (p *Player) Render() {
	rl.ImageResize(p.Image, int32(p.W), int32(p.H))
	rl.DrawTexture(p.Texture, int32(p.X), int32(p.Y), rl.White)
	// rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.W), int32(p.H), rl.Green);
}

func (p Player) GetPos() (float32, float32) {
	return p.X, p.Y
}

func (p *Player) SetPos(x float32, y float32) {
	p.X = x
	p.Y = y
}

func (p *Player) Move(gameAreas [][]ga.GameArea, window ga.Window, moveX float32, moveY float32) {
	if moveX != 0 && moveY != 0 {
			p.SetPos(p.X + float32(moveX) / math.Sqrt2, p.Y + float32(moveY) / math.Sqrt2)
		} else {
			p.SetPos(p.X + float32(moveX), p.Y + float32(moveY))
		}

	if (p.X < 0) {
			if (p.AreaX - 1 >= 0 && gameAreas[p.AreaY][p.AreaX - 1].IsActive) {
				p.AreaX -= 1;
				p.X = float32(window.Width - p.W)
			} else {
				p.SetPos(0, p.Y)
			}
		} else if (p.AreaX + 1 > 0 && p.X + float32(p.W) > float32(window.Width)) {
			if (gameAreas[p.AreaY][p.AreaX + 1].IsActive) {
				p.AreaX += 1
				p.X = 0
			} else {
				p.X = float32(window.Width) - float32(p.W)
			}
		}
		if (p.Y < 0) {
			if (p.AreaY - 1 >= 0 && gameAreas[p.AreaY - 1][p.AreaX].IsActive) {
				p.AreaY -= 1;
				p.Y = float32(window.Height - p.H)
			} else {
				p.SetPos(p.X, 0)
			}
		} else if (p.Y + float32(p.H) > float32(window.Height)) {
			if (p.AreaY + 1 > 0 && gameAreas[p.AreaY + 1][p.AreaX].IsActive) {
				p.AreaY += 1
				p.Y = 0
			} else {
				p.Y = float32(window.Height) - float32(p.H)
			}
		}
}