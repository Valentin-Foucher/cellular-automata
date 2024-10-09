package common

type Display[E comparable] interface {
	ShowGrid(*TwoDimensionalGrid[E])
	EraseGrid()
}
