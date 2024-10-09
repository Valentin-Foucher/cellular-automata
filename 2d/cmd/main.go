package main

import (
	core2d "cellular-automation/2d/core"
	"cellular-automation/common"
)

func main() {
	conf, err := common.ReadConfiguration[core2d.Configuration]("2d/configs/default.yaml")
	if err != nil {
		panic(err)
	}
	game, err := core2d.NewGame(conf)
	if err != nil {
		panic(err)
	}

	game.Start()
}
