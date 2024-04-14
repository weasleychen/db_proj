package util

import "math/rand"

func RandomString(length int) string {
	result := make([]byte, 0)

	for i := 0; i < length; i++ {
		result = append(result, byte(rand.Int()%128))
	}

	return string(result)
}
