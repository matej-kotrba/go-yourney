package magic

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	s "github.com/matej-kotrba/go-testing/src/spells"
)

type SpellPattern struct {
	Condition func(vectors []rl.Vector2) bool
	OnPass func(castPosition s.Destination, finalPosition s.Destination)
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
			// fmt.Printf(" %v ", lastAngle)
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
		OnPass: func (castPos s.Destination, endPos s.Destination) {
			s.NewFireball(castPos, endPos)
		},
	},
	"other": {
		Condition: func(vectors []rl.Vector2) bool {
			lines := make([]rl.Vector2, 0)

			const toleranceAngleSingle = 40 * (math.Pi / 180)
			const toleranceAngleAvarage = 40 * (math.Pi / 180)

			endX := float32(0)
			endY := float32(0)

			var avgLineAngle float64 = 0
			var avgLintCount int16 = 0
			var lastAngle float64 = 0

			var newLineStartIndex = 0
			var newLineXTotal float32 = 0
			var newLineYTotal float32 = 0

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
						lineX := float32(newLineXTotal) / (float32(i) - float32(newLineStartIndex))
						lineY := float32(newLineYTotal) / (float32(i) - float32(newLineStartIndex))

						fmt.Printf("X %v, Y %v \n", lineX, lineY)

						if (math.IsNaN(float64(lineX)) || math.IsInf(float64(lineX), 0)) {
							lineX = 0
						}

						if (math.IsNaN(float64(lineY)) || math.IsInf(float64(lineY), 0)) {
							lineY = 0
						}

						if (lineX != 0 || lineY != 0) {
							lines = append(lines, rl.Vector2{X: lineX,
								Y: lineY})
						}

						newLineStartIndex = i + 1
						newLineXTotal = 0
						newLineYTotal = 0
					}
				}

				newLineXTotal += vectors[i].X
				newLineYTotal += vectors[i].Y

				lastAngle = angle
				avgLineAngle += angle
			}
			
			// if (math.Abs(avgLineAngle / float64(avgLintCount)) > toleranceAngleAvarage / 2 + 90 * math.Pi / 180 || math.Abs(avgLineAngle / float64(avgLintCount)) < -toleranceAngleAvarage / 2 + 90 * math.Pi / 180) {
			// 	return false
			// }

			// if (math.Abs(float64(endY)) > tresholdY) {
			// 	return false
			// }

			// if (math.Abs(float64(endX)) > tresholdMaxX && math.Abs(float64(endX)) < tresholdMinX) {
			// 	return false
			// }

			fmt.Printf(" %v ", lines)

			return true
		},
		OnPass: func (castPos s.Destination, endPos s.Destination) {
			s.NewFireball(castPos, endPos)
		},
	},
}

func MatchSpellPattern(vectors []rl.Vector2, castPos s.Destination, endPos s.Destination) {
	for _, v := range Patterns {
		if (v.Condition(vectors)) {
			v.OnPass(castPos, endPos)
		}
	}

}