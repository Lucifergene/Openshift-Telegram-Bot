package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Lucifergene/openshift-telegram-bot/bot/utils"
)

func Start(update tgbotapi.Update) tgbotapi.MessageConfig {
	startStr := fmt.Sprintf("👋 <b>Hi %s, Welcome to Openshift Telegram Bot!</b> 🤖\n\nI can help you to deploy your application to OpenShift cluster. Just send me your <b>container image</b> name and I'll do the rest! 🚀\n\nFirst, lets login to your OpenShift cluster with /login command. 🔐", update.Message.From.FirstName)

	return utils.SendMsg(startStr, update)
}
