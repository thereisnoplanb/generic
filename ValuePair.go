package generic

// Represents a pair of values.
//
// Type1 - The type of the pair's first component.
// Type2 - The type of the pair's second component.
type ValuePair[Type1 any, Type2 any] struct {
	Item1 Type1
	Item2 Type2
}
