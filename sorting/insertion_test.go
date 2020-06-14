package sorting_test

import (
	"github.com/ameniGa/go-algorithms/sorting"
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	for _, tc := range sorting.CreateTTSort() {
		t.Run(tc.Name, func(t *testing.T) {
			sortable := sorting.Sortable{
				List:    tc.ToSort,
				Compare: tc.Compare,
			}
			got, err := sorting.Sort(sorting.Insertion, &sortable)
			if err != nil && !tc.HasError {
				t.Errorf("Expected success got error:%v", err)
			}
			if err == nil && tc.HasError {
				t.Errorf("Expected error got nil")
			}
			for idx, elem := range got {
				if reflect.ValueOf(elem).Interface() != reflect.ValueOf(tc.Expected).Index(idx).Interface() {
					t.Errorf("Expected:%v got:%v", tc.Expected, got)
				}
			}
		})
	}
}
