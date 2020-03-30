package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chippydip/go-sc2ai/api"
	"github.com/chippydip/go-sc2ai/botutil"
	"github.com/chippydip/go-sc2ai/client"
	"github.com/chippydip/go-sc2ai/runner"
)

const replayDir = "/Users/brandonbrombacher/Desktop/Rogue/01 zest first/"

type bot struct {
	*botutil.Bot

	myStartLocation    api.Point2D
	myNaturalLocation  api.Point2D
	enemyStartLocation api.Point2D

	camera api.Point2D
}

func main() {
	// Play a random map against a medium difficulty computer
	runner.SetComputer(api.Race_Random, api.Difficulty_Easy, api.AIBuild_RandomBuild)

	files, err := ioutil.ReadDir(replayDir)
	if err != nil {
		log.Println("Read directory error:", err)
	}
	replayPath := fmt.Sprintf("%v%v", replayDir, files[0].Name())
	runner.SetReplayPath(replayPath)
	runner.SetRealtime()

	// Create the agent and then start the game
	botutil.SetGameVersion()

	agent := client.AgentFunc(runAgent)

	runner.RunAgent(client.NewParticipant(api.Race_Protoss, agent, "StubBot"))
}

func runAgent(info client.AgentInfo) {

	bot := bot{Bot: botutil.NewBot(info)}
	bot.LogActionErrors()

	bot.init()
	for bot.IsInGame() {
		bot.doSmt()

		if err := bot.Step(1); err != nil {
			log.Print(err)
			break
		}
	}
}

func (bot *bot) init() {
	// Send a friendly hello
	bot.Chat("(glhf)")
}

func (bot *bot) doSmt() {

}
