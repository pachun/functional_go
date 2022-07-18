package functional_go

func Map[T any](items []T, function func(item T) T) []T {
	mapped := []T{}
	for _, item := range items {
		mapped = append(mapped, function(item))
	}
	return mapped
}
