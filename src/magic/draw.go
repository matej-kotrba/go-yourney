package magic

import rl "github.com/gen2brain/raylib-go/raylib"

type DrawPoint struct {
		X int16
		Y int16
}

type Draw struct {
	IsDrawing bool
	DrawedPattern []DrawPoint
	Color rl.Color
}