package common

type Automaton[E comparable] interface {
	NextState(*TwoDimensionalGrid[E]) *TwoDimensionalGrid[E]
}
