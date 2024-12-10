package batteries

func UniqueElements[T comparable](input []T) []T {
	var unique []T
	m := map[T]bool{}

	for _, v := range input {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
