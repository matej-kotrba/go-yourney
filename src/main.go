package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	g "github.com/matej-kotrba/go-testing/src/game"
	m "github.com/matej-kotrba/go-testing/src/magic"
	p "github.com/matej-kotrba/go-testing/src/player"
	s "github.com/matej-kotrba/go-testing/src/spells"
)

var window g.Window = g.Window{
	Width:  g.WINDOW_WIDTH,
	Height: g.WINDOW_HEIGHT,
}

const MOVE_SPEED = 10

var moveX int16 = 0;
var moveY int16 = 0;

func main() {
	rl.InitWindow(int32(window.Width), int32(window.Height), "Fun with go")
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
	player.Image = rl.LoadImage("static/imgs/seal-king.png")
	rl.ImageResize(player.Image, int32(player.W), int32(player.H))
	player.Texture = rl.LoadTextureFromImage(player.Image)

	var draw = m.Draw{
		Color: rl.NewColor(255, 255, 255, 255),
	}

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		
		// On key down
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

		// On mouse down
		if (rl.IsMouseButtonDown(rl.MouseLeftButton)) {
			if (!draw.IsDrawing) {
				draw.IsDrawing = true
				draw.DrawedPattern = append(draw.DrawedPattern, rl.NewVector2(float32(rl.GetMouseX()), float32(rl.GetMouseY())))
			}
			if (draw.IsDrawing) {	
				draw.DrawedPattern = append(draw.DrawedPattern, rl.NewVector2(float32(rl.GetMouseX()), float32(rl.GetMouseY())))
			}
		}

		// On mouse up
		if (rl.IsMouseButtonReleased(rl.MouseLeftButton)) {
			if (draw.IsDrawing) {
				dest := rl.GetMousePosition()
				draw.IsDrawing = false
				vectors := draw.GetVectors()
				draw.UsePattern(vectors, s.Destination{X: player.X, Y: player.Y}, s.Destination{X: dest.X, Y: dest.Y})
				
			}
		}

		player.Move(gameAreas, window, float32(moveX), float32(moveY))

		shouldRemoveProjectileIndexes := make([]int, 0, 10)

		for i, v := range s.Projectiles {
			if (v != nil) {
				v.Move()
				if (v.ShouldBeDeleted()) {
					shouldRemoveProjectileIndexes = append(shouldRemoveProjectileIndexes, i)
				}
			}
		}

		for _, v := range shouldRemoveProjectileIndexes {
			s.Projectiles = append(s.Projectiles[:v], s.Projectiles[v+1:]...)
		}

		moveX = 0
		moveY = 0

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Rendering process
		draw.DrawPattern()

		player.Render()
	
		for _, v := range s.Projectiles {
			if (v != nil) {
				v.Render()
			}
		}

		text := fmt.Sprintf("Area x: %v y: %v", player.AreaX, player.AreaY)

		rl.DrawText(text, 10, 10, 20, rl.White);

		rl.EndDrawing()
	}
}