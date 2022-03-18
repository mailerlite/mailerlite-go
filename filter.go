package mailerlite

// Filter is one of the arguments which has a name and a value
type Filter struct {
	// Name is the name of the field.
	Name string
	// Value is the value which the entry should be filtered by.
	Value interface{}
}

// NewFilter returns a new filter initialized with the given
// name and value.
func NewFilter(name string, value interface{}) *Filter {
	return &Filter{
		Name:  name,
		Value: value,
	}
}
