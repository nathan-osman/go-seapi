package seapi

// Value represents a piece of data retrieved from the API. The methods of this
// type enable safe access to the data in a map.
type Value map[string]interface{}

func (v Value) field(name string) interface{} {
	d, _ := v[name]
	return d
}

// List retrieves an array of items in the value or an empty array if the field
// is not present.
func (v Value) List(name string) []Value {
	vItems := make([]Value, 0)
	if iItems, ok := v.field(name).([]interface{}); ok {
		for _, iItem := range iItems {
			if v, ok := iItem.(map[string]interface{}); ok {
				vItems = append(vItems, Value(v))
			}
		}
	}
	return vItems
}

// String retrieves the string value with the specified name or an empty string
// if it does not exist.
func (v Value) String(name string) string {
	s, _ := v.field(name).(string)
	return s
}
