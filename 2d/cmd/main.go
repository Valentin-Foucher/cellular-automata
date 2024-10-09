package main

import (
	config2d "cellular-automation/2d/configs"
	game2d "cellular-automation/2d/game"
	"cellular-automation/common"
)

func main() {
	conf, err := common.ReadConfiguration[config2d.Configuration]("2d/configs/default.yaml")
	if err != nil {
		panic(err)
	}
	game, err := game2d.NewGame(conf)
	if err != nil {
		panic(err)
	}

	game.Start()
}
