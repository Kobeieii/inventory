package utils

func Map[T any, R any](array []T, f func(T) R) []R {
	result := make([]R, len(array))
	for i, x := range array {
		result[i] = f(x)
	}
	return result
}