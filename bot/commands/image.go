package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Lucifergene/openshift-telegram-bot/bot/utils"
	"github.com/Lucifergene/openshift-telegram-bot/pkg/deploy"
	"github.com/Lucifergene/openshift-telegram-bot/pkg/login"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Image(isLoggedIn *bool, defaultNamespace *string, update tgbotapi.Update) tgbotapi.Chattable {
	if *isLoggedIn == true {
		args := update.Message.CommandArguments()
		if args == "" {
			return utils.SendMsg("Please enter the container image name, namespace and port number in the following format:\n\n<code>/image image-name|port</code>", update)
		} else {
			values := strings.Split(args, "|")
			if len(values) != 2 {
				return utils.SendMsg("Please enter the container image name, namespace and port number in the following format:\n\n<code>/image image-name|port</code>", update)
			} else {
				imageName := strings.TrimSpace(values[0])
				port, err := strconv.Atoi(strings.TrimSpace(values[1]))
				if err != nil {
					return utils.SendMsg("Invalid port number. ❌🔢", update)
				}
				log.Printf("Image Name: %s | Port: %d\n", imageName, port)

				clientSet, _ := login.GetClientSet()
				if err != nil {
					return utils.SendMsg("<b>Connection Failed</b> ❌ Please try again. 😔", update)
				}

				_, err = deploy.DeployImage(clientSet, imageName, *defaultNamespace, port)
				if err != nil {
					log.Printf("Failed to deploy image: %v\n", err)
					return utils.SendMsg(fmt.Sprintf("Failed to deploy <b>%s</b> ❌ Please try again. 😔", imageName), update)
				} else {
					appURL, err := deploy.GetAppURL()
					if err != nil {
						log.Printf("Failed to get app URL: %v\n", err)
						return utils.SendMsg("Failed to fetch the Application URL. ❌🌐", update)
					} else {
						imageURL := "https://cataas.com/cat/says/APP%20DEPLOYED!!?height=200"
						message := "🎉 <b>Your application is deployed successfully!</b> 🚀✨\n\n<b>Note:</b> It may take a few minutes for the application to be available! 🕐🕑🕒🕓"
						replyMarkUp := tgbotapi.NewInlineKeyboardMarkup(
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonURL("🌐 Application URL", appURL),
							),
						)
						return utils.SendPhoto(imageURL, message, update, replyMarkUp)
					}
				}
			}
		}

	} else {
		return utils.SendMsg("Please login to your cluster first. 🔒🔑", update)
	}
}
