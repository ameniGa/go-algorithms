package sorting

import (
	"errors"
	"reflect"
)

type Algorithm int8

const (
	Insertion Algorithm = iota
)

type Sortable struct {
	List interface{}
	Compare
}

type sortable struct {
	list []interface{}
	Compare
}

// Sortable should be implemented by type to insertionSort
type Compare func(i, j interface{}) bool

func Sort(algorithm Algorithm, toSort *Sortable) ([]interface{}, error) {
	slice, err := getSlice(toSort)
	if err != nil {
		return nil, err
	}
	newSortable := sortable{
		list:    slice,
		Compare: toSort.Compare,
	}
	switch algorithm {
	case Insertion:
		insertionSort(&newSortable)
		return newSortable.list, nil
	default:
		return nil, errors.New("unknown algorithm")
	}
}

func getSlice(s *Sortable) ([]interface{}, error) {
	if s == nil || s.List == nil || s.Compare == nil {
		return nil, errors.New("invalid input")
	}
	value := reflect.ValueOf(s.List)
	if value.Kind() != reflect.Slice {
		return nil, errors.New("list should be a slice/array of any type")
	}
	length := value.Len()
	out := make([]interface{}, length)
	for i := 0; i < length; i++ {
		out[i] = value.Index(i).Interface()
	}
	return out, nil
}
