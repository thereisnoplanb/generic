package generic

// Represents the method that compares two objects of the same type.
//
// Parameters
//
//	first TObject - The first object to compare.
//	second TObject - The second object to compare.
//
// Return Value
//
//	result int - A signed integer that indicates the relative values of first and second, as shown in the following table.
//
//	+-----------------------+-------------------------------+
//	| Result                | Meaning                       |
//	+-----------------------+-------------------------------+
//	| Less than 0           | first is less than second.    |
//	| 0                     | first equals second.          |
//	| Greater than 0        | first is greater than second. |
//	+-----------------------+-------------------------------+
type Comparison[TObject any] func(first TObject, second TObject) (reslut int)
