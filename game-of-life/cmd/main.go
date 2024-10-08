package main

import (
	"cellular-automation/common"
	gameoflife "cellular-automation/game-of-life/core"
)

func main() {
	conf, err := common.ReadConfiguration[gameoflife.Configuration]("game-of-life/configs/default.yaml")
	if err != nil {
		panic(err)
	}
	game, err := gameoflife.NewGame(conf)
	if err != nil {
		panic(err)
	}

	game.Start()
}
