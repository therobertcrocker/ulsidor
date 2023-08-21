package quests

var ExpPerLevel = map[int]int{
	-2: 200,
	-1: 400,
	0:  500,
	1:  600,
	2:  800,
	3:  1000,
}

var ExpNeeded = 1000

func MissionsPerLevel(relativeLevel int) int {
	if relativeLevel < -2 {
		return ExpNeeded / ExpPerLevel[-2]
	}
	return ExpNeeded / ExpPerLevel[relativeLevel]
}

var ClassMultiplier = map[string]int{
	"hunt":        2,
	"acquisition": -1,
	"whisper":     1,
	"Knowledge":   0,
}
var GoldByLevel = map[int]int{
	1:  13,
	2:  23,
	3:  40,
	4:  67,
	5:  107,
	6:  167,
	7:  240,
	8:  333,
	9:  467,
	10: 667,
	11: 933,
	12: 1333,
	13: 2000,
	14: 3000,
	15: 4333,
	16: 6667,
	17: 10000,
	18: 16000,
	19: 26667,
	20: 46667,
}
