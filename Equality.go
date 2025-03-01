package generic

// Represents the method that reports two objects of the same type are equal.
//
// # Parameters
//
//	x TObject
//
// The first object to compare for equality.
//
//	y TObject
//
// The second object to compare for equality.
//
// # Returns
//
//	bool
//
// True if the specified objects are equal; otherwise, false.
type Equality[TObject any] func(x TObject, y TObject) bool
