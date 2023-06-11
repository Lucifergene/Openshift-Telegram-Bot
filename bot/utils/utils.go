package utils

import (
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMsg(msg string, update tgbotapi.Update, replyMarkUp ...interface{}) (message tgbotapi.MessageConfig) {
	message = tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	message.ParseMode = "html"
	if len(replyMarkUp) > 0 {
		message.ReplyMarkup = replyMarkUp[0]
	}
	return message
}

func SendPhoto(photoURL string, caption string, update tgbotapi.Update, replyMarkUp ...interface{}) (message tgbotapi.PhotoConfig) {
	message = tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(photoURL))
	message.Caption = caption
	message.ParseMode = "html"
	if len(replyMarkUp) > 0 {
		message.ReplyMarkup = replyMarkUp[0]
	}
	return message
}

func StartSession(duration *int, update tgbotapi.Update, timerActive *bool, globalLock *sync.Mutex, clusterURL *string, token *string, isLoggedIn *bool) {
	sessionTimeoutMins := time.Duration(*duration) * time.Minute
	if *timerActive {
		return
	}

	*timerActive = true
	time.AfterFunc(sessionTimeoutMins, func() {
		expireSession(update, globalLock, clusterURL, token, isLoggedIn)
	})
}

func expireSession(update tgbotapi.Update, globalLock *sync.Mutex, clusterURL *string, token *string, isLoggedIn *bool) {
	globalLock.Lock()
	defer globalLock.Unlock()

	*clusterURL = ""
	*token = ""
	*isLoggedIn = false

	SendMsg("Session expired! 🔒 Please /login again. 🔑", update)
}
