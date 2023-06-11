package bot

import (
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Lucifergene/openshift-telegram-bot/bot/commands"
	env "github.com/Lucifergene/openshift-telegram-bot/pkg/env"
)

var (
	clusterURL string
	token      string
	isLoggedIn bool

	defaultNamespace string
	defaultUser      string
	sessionTimeout   int
	globalLock       sync.Mutex
	timerActive      bool
)

func RunBot(config env.Config) {
	// setting the global vars
	clusterURL = ""
	token = ""
	isLoggedIn = false
	defaultNamespace = "default"
	defaultUser = "kube:admin"
	sessionTimeout = 10 /*minutes*/

	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				bot.Send(commands.Start(update))

			case "login":
				bot.Send(commands.Login(&isLoggedIn, update))

			case "default":
				bot.Send(commands.Default(update, &defaultNamespace, &defaultUser))

			case "connect":
				bot.Send(commands.Connect(&isLoggedIn, &clusterURL, &token, &defaultNamespace, &defaultUser, &sessionTimeout, &timerActive, &globalLock, update))

			case "image":
				bot.Send(commands.Image(&isLoggedIn, &defaultNamespace, update))

			case "logout":
				bot.Send(commands.Logout(&isLoggedIn, &clusterURL, &token, update))
			}

		} else {
			bot.Send(commands.GetLoginDetails(&clusterURL, &token, &defaultNamespace, &defaultUser, update))
		}

	}
}
