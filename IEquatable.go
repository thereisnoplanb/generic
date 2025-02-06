package generic

// Defines a generalized method that a value type or class implements to create a type-specific method for determining equality of instances.
//
// Type Parameters
//
//	TObject - The type of objects to compare.
type IEquatable[TObject any] interface {

	// Indicates whether the current object is equal to another object of the same type.
	//
	// Parameters
	//
	//	other TObject - An object to compare with current object.
	//
	// Return Value
	//
	//	result bool - true if the current object is equal to the other parameter; otherwise, false.
	Equal(other TObject) (result bool)
}
