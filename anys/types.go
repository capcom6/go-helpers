package anys

// AsPointer takes a value of any type and returns a pointer to that value.
func AsPointer[T any](v T) *T {
	return &v
}

// OrDefault returns the default value if the given pointer is nil, otherwise it returns the value at the pointer.
func OrDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

// ZeroDefault returns the default value if the given value is zero, otherwise it returns the value.
func ZeroDefault[T comparable](v T, def T) T {
	zero := new(T)
	if v == *zero {
		return def
	}
	return v
}
