package generic

// Represents a pair of values.
//
// TFirst - The type of the pair's first component.
// TSecond - The type of the pair's second component.
type ValuePair[TFirst any, TSecond any] struct {
	Item1 TFirst
	Item2 TSecond
}
