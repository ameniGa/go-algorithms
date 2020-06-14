package sorting

type TTSort struct {
	Name     string
	ToSort   interface{}
	Compare  Compare
	Expected interface{}
	HasError bool
}

// createTTSort creates table test for sort function
func CreateTTSort() []TTSort {
	type person struct {
		name string
		age  int
	}
	return []TTSort{
		{
			Name:   "already sorted list",
			ToSort: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Compare: func(i, j interface{}) bool {
				return i.(int) > j.(int)
			},
			Expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			Name:   "unsorted list",
			ToSort: []int{0, 10, 2, 11, 13, 5, 80, 24, 1, 0, 13},
			Compare: func(i, j interface{}) bool {
				return i.(int) > j.(int)
			},
			Expected: []int{0, 0, 1, 2, 5, 10, 11, 13, 13, 24, 80},
		},
		{
			Name:   "unsorted list of type string",
			ToSort: []string{"olfa", "amani", "fatma"},
			Compare: func(i, j interface{}) bool {
				return i.(string) > j.(string)
			},
			Expected: []string{"amani", "fatma", "olfa"},
		},
		{
			Name: "unsorted list of type struct: insertionSort by age",
			ToSort: []person{
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
			},
			Compare: func(i, j interface{}) bool {
				return i.(person).age > j.(person).age
			},
			Expected: []person{
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
			},
		},
		{
			Name:   "nil list",
			ToSort: nil,
			Compare: func(i, j interface{}) bool {
				return i.(int) > j.(int)
			},
			HasError: true,
		},
		{
			Name:     "missing compare func",
			ToSort:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Compare:  nil,
			HasError: true,
		},
		{
			Name:   "invalid input list",
			ToSort: "blabla",
			Compare: func(i, j interface{}) bool {
				return i.(string) > j.(string)
			},
			HasError: true,
		},
	}
}
