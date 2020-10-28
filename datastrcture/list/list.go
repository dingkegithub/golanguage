package list

type List interface {
	// insert node
	Insert(item interface{}) bool

	// delete node
	Delete(data interface{}) (interface{}, error)

	// query node
	Query(data interface{}) (int, error)

	// position node
	Index(i int) interface{}

	// list size
	Size() int
}
