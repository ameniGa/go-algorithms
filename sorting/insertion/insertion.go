package insertion

func Sort(list []int) []int {
	sortedList := make([]int, len(list))
	for i, value := range list {
		sortedList = insert(list, i, value)
	}
	return sortedList
}

func insert(list []int, pos int, value int) []int {
	i := pos - 1
	for i >= 0 && list[i] > value {
		list[i+1] = list[i]
		i--
	}
	//sort.Sort()
	list[i+1] = value
	return list
}
