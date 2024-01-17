package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	g "github.com/matej-kotrba/go-testing/src/game"
	p "github.com/matej-kotrba/go-testing/src/player"
)

var window g.Window = g.Window{
	Width:  800,
	Height: 450,
}

const MOVE_SPEED = 10

var moveX int16 = 0;
var moveY int16 = 0;

func main() {
	var gameAreas = make([][]g.GameArea, 10)
	for i := 0; i < len(gameAreas); i++ {
		gameAreas[i] = make([]g.GameArea, 10)
		for k := 0; k < len(gameAreas[i]); k++ {
			gameAreas[i][k].SetArea(false)
		}
	}

	gameAreas[0][1].SetArea(true)
	gameAreas[1][1].SetArea(true)
	gameAreas[1][2].SetArea(true)

	player := new(p.Player)
	player.SetPos(100, 100)
	player.W = 50
	player.H = 50
	player.AreaX = 1
	player.AreaY = 1

	rl.InitWindow(int32(window.Width), int32(window.Height), "Fun with go")
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

		player.Move(gameAreas, window, float32(moveX), float32(moveY))

		moveX = 0
		moveY = 0

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(x), int32(y), int32(player.W), int32(player.H), rl.Green);

		text := fmt.Sprintf("Area x: %v y: %v", player.AreaX, player.AreaY)

		rl.DrawText(text, 10, 10, 20, rl.White);

		rl.EndDrawing()
	}
}