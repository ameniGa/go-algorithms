package sorting

// Sort sorts a list which implements the Sortable interface
func insertionSort(sortable *sortable) {
	pos := 1
	for pos < len(sortable.list) {
		insert(sortable, pos)
		pos++
	}
}

func insert(sortable *sortable, pos int) {
	i := pos - 1
	value := sortable.list[pos]
	for i >= 0 && sortable.Compare(sortable.list[i], value) {
		sortable.list[i+1] = sortable.list[i]
		i--
	}
	sortable.list[i+1] = value
}
