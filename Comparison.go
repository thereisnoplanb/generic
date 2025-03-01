package generic

// Represents the method that compares two objects of the same type.
//
// # Parameters
//
//	x TObject
//
// The first object to compare.
//
//	y TObject
//
// The second object to compare.
//
// # Returns
//
//	int
//
// A signed integer that indicates the relative values of x and y, as shown in the following table.
//
//	+-----------------------+----------------------+
//	| Result                | Meaning              |
//	+-----------------------+----------------------+
//	| Less than 0           | x is less than y.    |
//	| 0                     | x equals y.          |
//	| Greater than 0        | x is greater than y. |
//	+-----------------------+----------------------+
type Comparison[TObject any] func(x TObject, y TObject) int
