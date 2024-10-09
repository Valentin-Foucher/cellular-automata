package common

type World[C any, D any] struct {
	Content           C
	ParticularContent []D
	M                 int
	N                 int
}
