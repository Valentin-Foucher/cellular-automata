package display2d

import (
	grid2d "cellular-automation/2d/grid"
	"cellular-automation/common"
	"errors"
)

type Display = common.Display[grid2d.Grid]

func Get(displayType string) (common.Display[grid2d.Grid], error) {
	switch displayType {
	case "console":
		return &ConsoleDisplay{}, nil
	default:
		return nil, errors.New("invalid display configuraton")
	}
}
