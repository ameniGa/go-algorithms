package insertion_test

import (
	"github.com/ameniGa/go-algorithms/sorting/insertion"
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	t.Run("already sorted list", func(t *testing.T) {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		mockSortable := insertion.SortableMock{
			GetLen: func() int {
				return len(list)
			},
			GetCompare: func(i, j interface{}) bool {
				return i.(int) > j.(int)
			},
			GetTranspose: func(i int, j interface{}) {
				list[i] = j.(int)
			},
			GetElement: func(i int) interface{} {
				return list[i]
			},
		}
		insertion.Sort(&mockSortable)
		if !reflect.DeepEqual(list, expected) {
			t.Errorf("expected same list: %v", list)
		}
		t.Log(list)
	})

	t.Run("unsorted list", func(t *testing.T) {
		list := []int{0, 10, 2, 11, 13, 5, 80, 24, 1, 0, 13}
		expected := []int{0, 0, 1, 2, 5, 10, 11, 13, 13, 24, 80}
		mockSortable := insertion.SortableMock{
			GetLen: func() int {
				return len(list)
			},
			GetCompare: func(i, j interface{}) bool {
				return i.(int) > j.(int)
			},
			GetTranspose: func(i int, j interface{}) {
				list[i] = j.(int)
			},
			GetElement: func(i int) interface{} {
				return list[i]
			},
		}
		insertion.Sort(mockSortable)
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("expected %v:got: %v", expected, list)
		}
		t.Log(list)
	})

	t.Run("unsorted list of type string", func(t *testing.T) {
		list := []string{"olfa", "amani", "fatma"}
		expected := []string{"amani", "fatma", "olfa"}
		mockSortable := insertion.SortableMock{
			GetLen: func() int {
				return len(list)
			},
			GetCompare: func(i, j interface{}) bool {
				return i.(string) > j.(string)
			},
			GetTranspose: func(i int, j interface{}) {
				list[i] = j.(string)
			},
			GetElement: func(i int) interface{} {
				return list[i]
			},
		}
		insertion.Sort(mockSortable)
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("expected %v:got: %v", expected, list)
		}
		t.Log(list)
	})

	t.Run("unsorted list of type struct: sort by age", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		list := []person{
			{
				name: "olfa",
				age:  28,
			},
			{
				name: "fatma",
				age:  27,
			},
			{
				name: "amani",
				age:  26,
			},
		}
		expected := []person{
			{
				name: "amani",
				age:  26,
			},
			{
				name: "fatma",
				age:  27,
			},
			{
				name: "olfa",
				age:  28,
			},
		}
		mockSortable := insertion.SortableMock{
			GetLen: func() int {
				return len(list)
			},
			GetCompare: func(i, j interface{}) bool {
				return i.(person).age > j.(person).age
			},
			GetTranspose: func(i int, j interface{}) {
				list[i] = j.(person)
			},
			GetElement: func(i int) interface{} {
				return list[i]
			},
		}
		insertion.Sort(mockSortable)
		if !reflect.DeepEqual(expected, list) {
			t.Errorf("expected %v:got: %v", expected, list)
		}
		t.Log(list)
	})
}
