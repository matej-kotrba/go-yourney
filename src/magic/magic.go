package magic

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpellPattern struct {
	Condition func(vectors []rl.Vector2) bool
	OnPass func()
}

var Patterns = map[string]SpellPattern{
	// line down shape
	"projectile": {
		Condition: func(vectors []rl.Vector2) bool {
			const tresholdY = 20
			const tresholdMinX = 20
			const tresholdMaxX = 100
			const toleranceAngleSingle = 40 * (math.Pi / 180)
			const toleranceAngleAvarage = 40 * (math.Pi / 180)

			endX := float32(0)
			endY := float32(0)

			var avgLineAngle float64 = 0
			var avgLintCount int16 = 0
			var lastAngle float64 = 0
			fmt.Printf(" %v ", lastAngle)
			for i := 0; i < len(vectors); i++ {
				if (vectors[i].X == 0 && vectors[i].Y == 0) {
					continue
				}
				endX += vectors[i].X
				endY += vectors[i].Y

				avgLintCount += 1
				angle := math.Atan2(float64(vectors[i].Y), float64(vectors[i].X))

				if (i != 1) {

					// Change this later
					if (math.Abs(angle - lastAngle) > toleranceAngleSingle / 2) {
						return false
					}
					
					
				}
				lastAngle = angle
				avgLineAngle += angle
			}
			
			if (math.Abs(avgLineAngle / float64(avgLintCount)) > toleranceAngleAvarage / 2 + 90 * math.Pi / 180 || math.Abs(avgLineAngle / float64(avgLintCount)) < -toleranceAngleAvarage / 2 + 90 * math.Pi / 180) {
				return false
			}
			// if (math.Abs(float64(endY)) > tresholdY) {
			// 	return false
			// }

			// if (math.Abs(float64(endX)) > tresholdMaxX && math.Abs(float64(endX)) < tresholdMinX) {
			// 	return false
			// }

			return true
		},
		OnPass: func () {
			
		},
	},
}

func MatchSpellPattern(vectors []rl.Vector2) {
	for _, v := range Patterns {
		if (v.Condition(vectors)) {
		}
	}

}