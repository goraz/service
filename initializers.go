package service

//Initializer Store Initializer
type Initializer struct {
	fn    func(interface{})
	order float32
}

//Initializers is array of  Initializer for sorting
type Initializers []Initializer

func (slice Initializers) Len() int {
	return len(slice)
}

func (slice Initializers) Less(i, j int) bool {
	return slice[i].order < slice[j].order
}

func (slice Initializers) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
