package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"skillissue"

	mcbot "github.com/Tnze/go-mc/bot"
)

var (
	name  = flag.String("name", "", "The name of the bot")
	uuid  = flag.String("uuid", "", "The UUID of the bot")
	token = flag.String("token", "", "The token of the bot")
	host  = flag.String("host", "localhost", "The host of the server")
)

func main() {
	flag.Parse()

	if *name == "" || *uuid == "" || *token == "" {
		flag.Usage()
		return
	}

	config, err := skillissue.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error while loading config: %s", err)
	}

	rand.Seed(time.Now().UnixMicro())
	bot := skillissue.NewBot(config)
	bot.Client.Auth = mcbot.Auth{
		Name: *name,
		UUID: *uuid,
		AsTk: *token,
	}

	bot.Logger.Fatal(bot.Run(*host))
}
