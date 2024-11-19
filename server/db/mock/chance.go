package mock

import "github.com/ZeFort/chance"

var C *chance.Chance
var cIsSet = false
var cSeed int64

func NewChance(seed int64) {
	if !cIsSet || seed != cSeed {
		C = chance.NewS(seed)
		cIsSet = true
		cSeed = seed
	}
}
