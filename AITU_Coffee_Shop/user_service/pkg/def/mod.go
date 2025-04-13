package def

// Pointer return pointer on any value.
func Pointer[T any](v T) *T { return &v }
