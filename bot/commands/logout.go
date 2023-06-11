package commands

import (
	"github.com/Lucifergene/openshift-telegram-bot/bot/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Logout(isLoggedIn *bool, clusterURL, token *string, update tgbotapi.Update) tgbotapi.MessageConfig {
	if *isLoggedIn == true {
		*isLoggedIn = false
		*clusterURL = ""
		*token = ""
		// replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		// 	tgbotapi.NewInlineKeyboardRow(
		// 		tgbotapi.NewInlineKeyboardButtonURL("🔒 Request another token! 🔄", fmt.Sprintf("%s/oauth/token/request", clusterURL)),
		// 	),
		// )
		return utils.SendMsg("<b>🔓 Logged out successfully! 👋😊 </b>", update)
	} else {
		return utils.SendMsg("You are not logged in. 🚫🔒", update)
	}
}
