package generic

// Represents a function that defines a set of criteria and determines whether the specified object meets those criteria.
//
// # Parameters
//
//	object TObject
//
// The object to compare against the criteria defined within the method represented by this Predicate.
//
// # Returns
//
//	bool
//
// True if object meets the criteria defined within the method represented by this Predicate; otherwise, false.
type Predicate[TObject any] func(object TObject) bool
