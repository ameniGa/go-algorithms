package insertion_test

import (
	"github.com/ameniGa/go-algorithms/sorting/insertion"
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	t.Run("already sorted list", func(t *testing.T) {
		list := []int{0,1,2,3,4,5,6,7,8,9}
		sortedList := insertion.Sort(list)
		if !reflect.DeepEqual(list,sortedList) {
			t.Errorf("expected same list: %v",sortedList)
		}
		t.Log(sortedList)
	})
	t.Run("unsorted list", func(t *testing.T) {
		list := []int{0,10,2,11,13,5,80,24,1,0,13}
		expected := []int{0,0,1,2,5,10,11,13,13,24,80}
		sortedList := insertion.Sort(list)
		if !reflect.DeepEqual(expected,sortedList) {
			t.Errorf("expected same list: %v",sortedList)
		}
		t.Log(sortedList)
	})
}


