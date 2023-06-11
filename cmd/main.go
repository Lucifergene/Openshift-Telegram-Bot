package main

import (
	"log"

	env "github.com/Lucifergene/openshift-telegram-bot/pkg/env"

	"github.com/Lucifergene/openshift-telegram-bot/bot"
)

func main() {
	config, err := env.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	bot.RunBot(config)

	select {}
}
