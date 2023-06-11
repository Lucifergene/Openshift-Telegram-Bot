package commands

import (
	"strings"

	"github.com/Lucifergene/openshift-telegram-bot/bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Default(update tgbotapi.Update, defaultNamespace *string, defaultUser *string) tgbotapi.MessageConfig {
	args := update.Message.CommandArguments()
	if args == "" {
		return utils.SendMsg("Please enter the default namespace and user in the following format:\n\n<b>/default namespace|user</b>", update)
	} else {
		values := strings.Split(args, "|")
		if len(values) != 2 {
			return utils.SendMsg("Please enter the default namespace and user in the following format:\n\n<b>/default namespace|user</b>", update)
		} else {
			*defaultNamespace = strings.TrimSpace(values[0])
			*defaultUser = strings.TrimSpace(values[1])
			return utils.SendMsg("<b>Default values updated to the following:</b>\n\n<b>Namespace:</b> "+*defaultNamespace+"\n<b>User:</b> "+*defaultUser+"\n\n<b>Now you can connect to your cluster with the 🔗 /connect command. 🌐💻</b>", update)
		}
	}
}
