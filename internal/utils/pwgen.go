package utils

import (
	"math/rand/v2"
	"strings"
)

var dicts = [][]rune{[]rune("аоуыэяюие"), []rune("бвгджзклмнпрстфхч")}

// шанс на смену словаря увеличивается вплоть до 100% за указанное кол-во шагов
const dictChangeEvery = 3
const pwLength = 6
const runeSize = 4

func GeneratePassword() string {
	var pw strings.Builder
	pw.Grow(pwLength * runeSize)

	dictI := rand.IntN(len(dicts))
	lastDictChange := -1
	prevRune := struct {
		runeI int
		dictI int
	}{-1, -1}
	for i := 0; i < pwLength; i++ {
		dictI %= len(dicts)
		dict := dicts[dictI]

		var runeI int
		// исключение повтора рун
		if prevRune.dictI == dictI {
			runeI = rand.IntN(len(dict) - 1)

			if runeI >= prevRune.runeI {
				runeI++
			}
		} else {
			runeI = rand.IntN(len(dict))
		}

		prevRune.dictI = dictI
		prevRune.runeI = runeI

		pw.WriteRune(dict[runeI])

		noDictChangeSteps := i - lastDictChange
		dictChangeChance := dictChangeEvery - noDictChangeSteps
		if dictChangeChance <= 0 || rand.IntN(dictChangeChance) == 0 {
			dictI++
			lastDictChange = i
		}
	}

	return pw.String()
}
