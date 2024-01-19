package magic

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpellPattern struct {
	condition func(vectors []rl.Vector2) bool
}

type Spell struct {
	pattern SpellPattern
	castX   int16
	castY   int16
}

var Patterns = map[string]SpellPattern{
	// "v" shape
	"projectile": {
		condition: func(vectors []rl.Vector2) bool {
			const tresholdY = 20
			const tresholdMinX = 20
			const tresholdMaxX = 100
			
			endX := float32(0)
			endY := float32(0)

			for i := 0; i < len(vectors); i++ {
				endX += vectors[i].X
				endY += vectors[i].Y
			}

			if (math.Abs(float64(endY)) > tresholdY) {
				return false
			}

			// if (math.Abs(float64(endX)) > tresholdMaxX && math.Abs(float64(endX) < tresholdMinX)) {
			// 	return false
			// }
		},
	},
}

func EvaluateSpell([]struct {
	x int16
	y int16
}) {
}