package commands

import (
	"log"
	"sync"

	"github.com/Lucifergene/openshift-telegram-bot/bot/utils"
	"github.com/Lucifergene/openshift-telegram-bot/pkg/login"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Connect(isLoggedIn *bool, clusterURL *string, token *string, defaultNamespace *string, defaultUser *string, sessionTimeout *int, timerActive *bool, globalLock *sync.Mutex, update tgbotapi.Update) tgbotapi.MessageConfig {
	if *isLoggedIn == false {
		if *clusterURL != "" && *token != "" {
			err := login.ClusterLogin(*clusterURL, *token, *defaultNamespace, *defaultUser)
			var msg string
			if err != nil {
				log.Printf("Failed to login: %v\n", err)
				msg = "<b>❌ Failed to login. Please try again. 😔</b>"
				*clusterURL = ""
				*token = ""
			} else {
				*isLoggedIn = true
				*timerActive = false
				utils.StartSession(sessionTimeout, update, timerActive, globalLock, clusterURL, token, isLoggedIn)
				msg = "<b>🔓 You have successfully logged in to your cluster ✅🎉 !!!</b>\n\n<b>Now you can deploy your application with the /image command. 🤖</b>\n\nEnter the container image name and port number in the following format: <code>/image image-name|port</code>\n\n<b>Note:</b> For security reasons, your session will expire in 10 minutes. 🔒"
			}
			return utils.SendMsg(msg, update)
		} else {
			return utils.SendMsg("Please login to your cluster with the 🔒 /login command. 💻", update)
		}
	} else {
		return utils.SendMsg("A cluster is already connected! ✅🔗", update)
	}
}
