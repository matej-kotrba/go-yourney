package magic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Draw struct {
	IsDrawing bool
	DrawedPattern []rl.Vector2
	Color rl.Color
}

func (d *Draw) DrawPattern() {
	for i := 0; i < len(d.DrawedPattern) - 1; i++ {
		rl.DrawLineEx(d.DrawedPattern[i], d.DrawedPattern[i+1], 10, rl.White)
		rl.DrawCircle(int32(d.DrawedPattern[i].X), int32(d.DrawedPattern[i].Y), 5, rl.White)
	}
}

func (d *Draw) ClearPattern() {
	d.DrawedPattern = d.DrawedPattern[:0]
}

func (d *Draw) GetVectors() []rl.Vector2 {
	// Change later maybe
	tempSlice := make([]rl.Vector2, 0)
	for i := 0; i < len(d.DrawedPattern) - 1; i++ {
		tempSlice = append(tempSlice, rl.NewVector2(float32(d.DrawedPattern[i+1].X) - float32(d.DrawedPattern[i].X), float32(d.DrawedPattern[i+1].Y) - float32(d.DrawedPattern[i].Y)))
	}
	return tempSlice
}