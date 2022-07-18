package functional_go

func Map[T any, K any](items []T, function func(item T) K) []K {
	mapped := []K{}

	for _, item := range items {
		mapped = append(mapped, function(item))
	}

	return mapped
}

func Filter[T any](items []T, function func(item T) bool) []T {
	filtered := []T{}

	for _, item := range items {
		if function(item) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

func Reduce[T any, V any](items []V, initial T, function func(accumulator T, nextNumber V) T) T {
	reduced := initial

	for _, item := range items {
		reduced = function(reduced, item)
	}

	return reduced
}
