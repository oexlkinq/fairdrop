package main

import (
	"math/rand/v2"
	"strings"
)

var dicts = []string{"аоуыэяюие", "бвгджзклмнпрстфхч"}

// шанс на смену словаря увеличивается вплоть до 100% за указанное кол-во шагов
const dictChangeEvery = 2
const letterSize = 2
const pwLength = 6
const pwSize = pwLength * letterSize

func GeneratePassword() string {
	var pw strings.Builder
	pw.Grow(pwSize)

	dictI := rand.IntN(len(dicts))
	lastDictChange := 0
	for i := 0; i < pwLength; i++ {
		dictI %= len(dicts)
		dict := dicts[dictI]

		dictLetterI := rand.IntN(len(dict) / letterSize)
		dictByteI := dictLetterI * letterSize

		pw.WriteByte(dict[dictByteI])
		pw.WriteByte(dict[dictByteI+1])

		noDictChangeSteps := i - lastDictChange
		dictChangeChance := dictChangeEvery - noDictChangeSteps
		if dictChangeChance <= 0 || rand.IntN(dictChangeChance) == 0 {
			dictI++
			lastDictChange = i
		}
	}

	return pw.String()
}
