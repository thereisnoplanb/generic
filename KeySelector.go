package generic

// Represents a function that extracts a key from the specified object.
//
// # Parameters
//
//	object TObject
//
// The object to extract a key from.
//
// # Returns
//
//	TKey
//
// The key extracted from the object.
type KeySelector[TObject any, TKey comparable] func(object TObject) TKey
