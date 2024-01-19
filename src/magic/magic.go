package magic

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpellPattern struct {
	Condition func(vectors []rl.Vector2) bool
}

type Spell struct {
	pattern SpellPattern
	castX   int16
	castY   int16
}

var Patterns = map[string]SpellPattern{
	// "v" shape
	"projectile": {
		Condition: func(vectors []rl.Vector2) bool {
			const tresholdY = 20
			const tresholdMinX = 20
			const tresholdMaxX = 100
			const toleranceAngleSingle = 100 * (math.Pi / 180) 
			const toleranceAngleAvarage = 100 * (math.Pi / 180) 

			endX := float32(0)
			endY := float32(0)

			var avgLineAngle float64 = 0
			var avgLintCount int16 = 0
			var lastAngle float64 = 0

			for i := 0; i < len(vectors); i++ {
				if (vectors[i].X == 0 && vectors[i].Y == 0) {
					continue
				}
				endX += vectors[i].X
				endY += vectors[i].Y

				avgLintCount++
				angle := math.Atan2(float64(vectors[i].Y), float64(vectors[i].X))

				// Change this later
				if (math.Abs(angle - lastAngle) > toleranceAngleSingle / 2) {
					fmt.Printf("%v %v \n", angle, toleranceAngleSingle / 2)
					return false
				}

				if (math.Abs(avgLineAngle / float64(avgLintCount)) > toleranceAngleAvarage / 2) {
					return false
				}

				lastAngle = angle
				avgLineAngle += angle
			}

			// if (math.Abs(float64(endY)) > tresholdY) {
			// 	return false
			// }

			// if (math.Abs(float64(endX)) > tresholdMaxX && math.Abs(float64(endX)) < tresholdMinX) {
			// 	return false
			// }

			return true
		},
	},
}

func EvaluateSpell([]struct {
	x int16
	y int16
}) {
}