package batteries

// SplitSlice splits a slice evenly into N parts.
// If the slice can't be split evenly, the last split will be short
func SplitSlice[T any](input []T, splits int) [][]T {
	// inputLen := len(input)
	//danglingAmount := len(input) % splits
	splitInput := make([][]T, splits)

	for i, val := range input {
		// get split index
		splitSubindex := i % splits
		splitIndex := (i - splitSubindex) / splits

		splitInput[splitIndex][splitSubindex] = val
	}

	return splitInput
}
