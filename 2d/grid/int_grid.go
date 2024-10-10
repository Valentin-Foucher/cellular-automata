package grid2d

type IntGrid = BaseGrid[int, int]

func NewIntGrid(m, n int) *IntGrid {
	return NewBaseGrid[int, int](m, n)
}
