package generic

// Represents a pair of values.
//
//	TObject1
//
// The type of the pair's first component.
//
//	TObject2
//
// The type of the pair's second component.
type ValuePair[TObject1 any, TObject2 any] struct {
	Item1 TObject1
	Item2 TObject2
}
