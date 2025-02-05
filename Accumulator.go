package generic

// Represents the method that accumulates object to accumulator.
//
// Parameters
//
//	accumulator TAccumulator - The current value of the accumulator.
//	object TObject - The object to accumulate.
//
// Return Value
//
//	result TAccumulator - The accumulated value of the current value of the accumulator and the object to accumulate.
type Accumulator[TObject any, TAccumulator any] func(accumulator TAccumulator, object TObject) (result TAccumulator)
