package common

type Display[W any] interface {
	Print(W)
	Flush()
}
