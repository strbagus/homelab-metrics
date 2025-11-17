package utils

func ArrFilter[T any](arr []T, test func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, item := range arr {
		if test(item) {
			result = append(result, item)
		}
	}
	return result
}
