package generic

// Represents the method that reports two objects of the same type are equal.
//
// Parameters
//
//	first TObject - The first object to compare.
//	second TObject - The second object to compare.
//
// Return Value
//
//	result bool - true if the specified objects are equal; otherwise, false.
type Equality[TObject any] func(first TObject, second TObject) (result bool)
