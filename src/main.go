package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	p "github.com/matej-kotrba/go-testing/src/player"
)

const MOVE_SPEED = 5

var moveX int16 = 0;
var moveY int16 = 0;

func main() {
	player := new(p.Player)
	player.SetPos(100, 100)

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		x, y := player.GetPos()

		if (rl.IsKeyDown(rl.KeyRight)) {
			moveX += MOVE_SPEED;
		}
		if (rl.IsKeyDown(rl.KeyLeft)) {
			moveX -= MOVE_SPEED;
		}
		if (rl.IsKeyDown(rl.KeyDown)) {
			moveY += MOVE_SPEED;
		}
		if (rl.IsKeyDown(rl.KeyUp)) {
			moveY -= MOVE_SPEED;
		}

		if x != 0 && y != 0 {
			player.SetPos(x + float32(moveX) / math.Sqrt2, y + float32(moveY) / math.Sqrt2)
		} else {
			player.SetPos(x + float32(moveX), y + float32(moveY))
		}

		moveX = 0
		moveY = 0

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(x), int32(y), 50, 50, rl.Green);

		rl.EndDrawing()
	}
}