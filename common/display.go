package common

type Display[W any] interface {
	Show(W)
	Erase()
}
