package main

import (
	"fmt"
	"math/rand/v2"
)

type PasswordComplexity int

const (
	Low PasswordComplexity = iota
	Medium
	High
)

func generatePassword(complexity PasswordComplexity) string {
	return fmt.Sprintf("password%d%d", complexity, rand.Int())
}
