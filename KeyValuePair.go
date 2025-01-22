package generic

// Represents a pair of key and value.
//
// TKey - The type of the key.
// TValue - The type of the value.
type KeyValuePair[TKey comparable, TValue any] struct {
	Key   TKey
	Value TValue
}
