package util

func NewType[T any](value T) *T {
	p := new(T)
	*p = value
	return p
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func ConvertArrayType[T Number, P Number](arr []T) []P {
	var result []P
	for index := range arr {
		result[index] = P(arr[index])
	}
	return result
}
