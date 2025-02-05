package generic

// Represents a function that produces a value from the specified object.
//
// Parameters
//
//	object TObject - The object to produce a value from.
//
// Return Value
//
//	value TValue - The value produced from the object.
type ValueSelector[TObject any, TValue any] func(object TObject) (value TValue)
