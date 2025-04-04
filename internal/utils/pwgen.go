package utils

import (
	"math/rand/v2"
	"strings"
)

var dicts = []string{"аоуыэяюие", "бвгджзклмнпрстфхцчшщ"}

const letterSize = 2
const pwLength = 6
const pwSize = pwLength * letterSize

func GeneratePassword() string {
	var pw strings.Builder
	pw.Grow(pwSize)

	dictIShift := rand.IntN(2)
	for i := 0; i < pwLength; i++ {
		// побитовое И вернёт всегда либо 0 либо 1. больше не нужно, т.к. словаря 2
		dictI := (i + dictIShift) & 1
		dict := dicts[dictI]

		dictLetterI := rand.IntN(len(dict) >> 1)
		dictByteI := dictLetterI << 1

		pw.WriteByte(dict[dictByteI])
		pw.WriteByte(dict[dictByteI+1])
	}

	return pw.String()
}
