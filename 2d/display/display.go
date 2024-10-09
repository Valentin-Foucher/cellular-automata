package display2d

import (
	config2d "cellular-automation/2d/configs"
	grid2d "cellular-automation/2d/grid"
	"cellular-automation/common"
	"errors"
)

type Display = common.Display[grid2d.Grid]

func Get(conf *config2d.Configuration) (common.Display[grid2d.Grid], error) {
	switch conf.Display {
	case "console":
		return &ConsoleDisplay{}, nil
	case "gif":
		return &GifDisplay{delay: conf.Delay}, nil
	default:
		return nil, errors.New("invalid display configuraton")
	}
}
