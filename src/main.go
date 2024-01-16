package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	p "github.com/matej-kotrba/go-testing/src/player"
)

type Window struct {
	Width int16
	Height int16
}

var window Window = Window{
	Width: 800,
	Height: 450,
}

const MOVE_SPEED = 10

var moveX int16 = 0;
var moveY int16 = 0;

func main() {
	player := new(p.Player)
	player.SetPos(100, 100)
	player.W = 50
	player.H = 50

	rl.InitWindow(int32(window.Width), int32(window.Height), "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		x, y := player.GetPos()

		if (rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)) {
			moveX += MOVE_SPEED;
		}
		if (rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)) {
			moveX -= MOVE_SPEED;
		}
		if (rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS)) {
			moveY += MOVE_SPEED;
		}
		if (rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW)) {
			moveY -= MOVE_SPEED;
		}

		if x != 0 && y != 0 {
			player.SetPos(x + float32(moveX) / math.Sqrt2, y + float32(moveY) / math.Sqrt2)
		} else {
			player.SetPos(x + float32(moveX), y + float32(moveY))
		}

		newX, newY := player.GetPos()

		if (newX < 0) {
			player.SetPos(0, player.Y)
		} else if (newX + float32(player.W) > float32(window.Width)) {
			player.X = float32(window.Width) - float32(player.W)
		}
		if (newY < 0) {
			player.SetPos(player.X, 0)
		} else if (newY + float32(player.H) > float32(window.Height)) {
			player.Y = float32(window.Height) - float32(player.H)
		}

		moveX = 0
		moveY = 0

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(x), int32(y), int32(player.W), int32(player.H), rl.Green);

		rl.EndDrawing()
	}
}