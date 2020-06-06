package insertion

type Sortable interface {
	Len() int
	Compare(i, j interface{}) bool
	Transpose(i int, j interface{})
	Get(i int) interface{}
}

func Sort(list Sortable) {
	pos := 1
	for pos < list.Len() {
		insert(list, pos)
		pos++
	}
}

func insert(list Sortable, pos int) {
	i := pos - 1
	value := list.Get(pos)
	for i >= 0 && list.Compare(list.Get(i), value) {
		list.Transpose(i+1, list.Get(i))
		i--
	}
	list.Transpose(i+1, value)
}
