package commands

import (
	"github.com/Lucifergene/openshift-telegram-bot/bot/utils"
	"github.com/Lucifergene/openshift-telegram-bot/pkg/login"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Login(isLoggedIn *bool, update tgbotapi.Update) tgbotapi.MessageConfig {
	if *isLoggedIn == false {
		return utils.SendMsg("<b>Please enter the Openshift Cluster URL: 🔗</b>", update)
	} else {
		return utils.SendMsg("You are already logged in! 👍😊", update)
	}
}

func GetLoginDetails(clusterURL, token, defaultNamespace, defaultUser *string, update tgbotapi.Update) tgbotapi.MessageConfig {
	if *clusterURL == "" {
		if login.IsUrl(update.Message.Text) == false {
			return utils.SendMsg("Please enter a valid URL. 🔗", update)
		} else {
			*clusterURL = update.Message.Text
			loginStr := "<b>Please follow these steps to fetch yout API Token:</b>\n\n1. Open the given Authorization URL 🔗\n2. Sign in with your OpenShift credentials 🔐\n3. Click on the <b>Display Token</b> link 🔒\n4. Copy the generated API token and paste it here 📋"
			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL("Authorization URL", login.GetAuthURL(*clusterURL)),
				),
			)
			return utils.SendMsg(loginStr, update, inlineKeyboard)
		}

	} else if *token == "" {
		*token = update.Message.Text

		return utils.SendMsg("<b>Default values are set to the following:</b>\n\n<b>Namespace:</b> "+*defaultNamespace+"\n<b>User:</b> "+*defaultUser+"\n\n<b>You can change these values with /default command</b>\n<b>Format:</b> <code>/default namespace|user</code>\n\n<b>Now you can connect to your cluster with the 🔗 /connect command. 🌐💻</b>", update)

	}
	return utils.SendMsg("Please use the commands from the Menu. 📋", update)
}
