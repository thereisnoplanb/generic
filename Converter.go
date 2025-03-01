package generic

// Represents the method that converts one object to another.
//
// # Parameters
//
//	object TObject
//
// The object to convert.
//
// # Returns
//
//	TResult
//
// Object converted from input object.
type Converter[TObject any, TResult any] func(object TObject) TResult
