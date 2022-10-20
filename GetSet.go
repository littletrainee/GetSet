// Encapsulation for any value
package GetSet

// value can get and set
type Type[T any] struct{ value T }

// Get Value
func (t *Type[T]) Get() T { return t.value }

// Set Value
func (t *Type[T]) Set(value T) { t.value = value }
