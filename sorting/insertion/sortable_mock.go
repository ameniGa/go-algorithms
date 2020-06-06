package insertion

type SortableMock struct {
	GetLen       func() int
	GetCompare   func(i, j interface{}) bool
	GetTranspose func(i int, j interface{})
	GetElement   func(i int) interface{}
}

func (s SortableMock) Len() int {
	return s.GetLen()
}

func (s SortableMock) Compare(i, j interface{}) bool {
	return s.GetCompare(i, j)
}

func (s SortableMock) Transpose(i int, j interface{}) {
	s.GetTranspose(i, j)
}

func (s SortableMock) Get(i int) interface{} {
	return s.GetElement(i)
}
