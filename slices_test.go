package batteries

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	got := SplitSlice[int]([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3)

	if !reflect.DeepEqual(got, [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}) {
		t.Errorf("got %v, wanted %v", got, [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}})
	}

}
