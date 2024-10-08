package gameoflife

import "cellular-automation/common"

type Grid = common.Grid[bool]

func NewGrid(m, n int) *Grid {
	return common.NewGrid(m, n, false)
}
