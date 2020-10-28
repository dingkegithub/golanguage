package list

// compare data
type CompareFunc func(interface{}, interface{}) (int, error)

// list option config
type Option func(*ListOption)

// list options
type ListOption struct {
	Compare CompareFunc
}

func (lo *ListOption) Apply(opts ...Option) {
	for _, o := range opts {
		o(lo)
	}
}

func WithIntCompare() Option {
	return func(option *ListOption) {
		option.Compare = defaultIntCompare
	}
}

func defaultIntCompare(a interface{}, b interface{}) (int, error) {

	aI, ok := a.(int)
	if !ok {
		return 0, ErrorTypeInt
	}

	bI, ok := b.(int)
	if !ok {
		return 0, ErrorTypeInt
	}

	if aI == bI {
		return 0, nil
	} else if aI > bI {
		return 1, nil
	} else {
		return -1, nil
	}
}
