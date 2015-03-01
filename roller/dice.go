package roller

import (
	"math/rand"
	"time"
)

func DoRolls(spec RollSpec) RollResults {
	var localResults RollResults

	rand.Seed(time.Now().Unix())
	var eachSet int64
	for eachSet = 0; eachSet < spec.Times; eachSet++ {
		localResults.Rolls = append(localResults.Rolls, roll(spec))
		localResults.Count++
	}

	return localResults
}

func roll(spec RollSpec) SetResult {
	var result SetResult

	var eachDie int64
	var sidesInt int = int(spec.Sides)
	for eachDie = 0; eachDie < spec.DieCount; eachDie++ {
		die := rand.Intn(sidesInt) + 1
		result.Total += die
		result.Dies = append(result.Dies, die)
	}

	return result
}
