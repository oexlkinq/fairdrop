package utils

import (
	"fmt"
	"math/rand/v2"
)

func GeneratePassword() string {
	// TODO: реализовать генерацию пароля
	// TODO: реализовать проверку уникальности пароля по базе перед возвратом
	return fmt.Sprintf("pw%d", rand.Int()%10000)
}
