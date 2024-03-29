package spells

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	g "github.com/matej-kotrba/go-testing/src/game"
)

const speed = 5

type Fireball struct {
	X float32
	Y  float32
	SX float32
	SY float32
	R int16
	image rl.Image
	texture rl.Texture2D
}

func NewFireball(spawnDest Destination, dest Destination) {
	f := new(Fireball)

	f.R = 30

	f.image = *rl.LoadImage("static/imgs/spells/fireball.png")
	rl.ImageResize(&f.image, int32(f.R) * 2, int32(f.R) * 2)
	f.texture = rl.LoadTextureFromImage(&f.image)

	xLine := dest.X - spawnDest.X
	yLine := dest.Y - spawnDest.Y
	c := float32(math.Sqrt(math.Pow(float64(xLine), 2) + math.Pow(float64(yLine), 2)))

	f.Y  = spawnDest.Y
	f.X  = spawnDest.X
	f.SX = xLine / c
	f.SY = yLine / c
	
	Projectiles = append(Projectiles, f)
}

func (f *Fireball) Move() {
	f.X += f.SX * speed
	f.Y += f.SY * speed
}

func (f *Fireball) ShouldBeDeleted() bool {
	return f.X - float32(f.R) >= g.WINDOW_WIDTH || f.X + float32(f.R) <= 0 || f.Y - float32(f.R) >= g.WINDOW_HEIGHT || f.Y + float32(f.R) <= 0
}

func (f *Fireball) Render() {
	rl.ImageResize(&f.image, int32(f.R * 2), int32(f.R * 2))
	rl.DrawTexture(f.texture, int32(f.X), int32(f.Y), rl.White)
}