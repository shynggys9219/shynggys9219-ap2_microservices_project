package def

// Pointer return pointer on any value.
func Pointer[T any](v T) *T { return &v }

// IsPointerValueChanged checks if two pointer values are different
func IsPointerValueChanged[T comparable](oldPtr, newPtr *T) bool {
	return (oldPtr == nil) != (newPtr == nil) || (oldPtr != nil && *oldPtr != *newPtr)
}
