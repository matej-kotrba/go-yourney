package magic

type SpellPattern struct {
	grid [9][9]bool
}

type Spell struct {
	pattern SpellPattern
	castX   int16
	castY   int16
}

var Patterns = map[string]Spell{
	"projectile": {
		pattern: SpellPattern{
			grid: [9][9]bool{
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
				{false, false, false, false, true, false, false, false, false},
			},
		},
	},
}

func EvaluateSpell([]struct {
	x int16
	y int16
}) {
}