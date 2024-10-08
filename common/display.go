package common

type Display[E comparable] interface {
	ShowGrid(*Grid[E])
	EraseGrid(*Grid[E])
}
