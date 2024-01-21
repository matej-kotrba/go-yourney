package spells

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const speed = 5

type Fireball struct {
	X float32
	Y  float32
	SX float32
	SY float32
	W int16
	H int16
	image rl.Image
	texture rl.Texture2D
}

func (f *Fireball) New(spawnDest Destination, dest Destination) {
	f.image = *rl.LoadImage("static/imgs/spells/fireball.png")
	rl.ImageResize(&f.image, int32(f.W), int32(f.H))
	f.texture = rl.LoadTextureFromImage(&f.image)

	xLine := spawnDest.X - dest.X
	yLine := spawnDest.Y - dest.Y
	c := float32(math.Sqrt(math.Exp2(float64(xLine)) + math.Exp2(float64(yLine))))

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

func (f *Fireball) Render() {
	rl.ImageResize(&f.image, int32(f.W), int32(f.H))
	rl.DrawTexture(f.texture, int32(f.X), int32(f.Y), rl.White)
}