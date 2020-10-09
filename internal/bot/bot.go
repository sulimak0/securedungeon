package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sulimak0/securedungeon/internal/models"
	"log"
	"net/http"
	"time"
)

type Unit interface {
	TurnOn()
	TurnOff()
	Url() string
	State() bool
	GetMessage() models.Alarm
}

type Bot struct {
	bot     *tgbotapi.BotAPI
	users   []int64
	offset  int
	timeout int
	unit    Unit
}

func NewBot(token string, offset, timeout int, unit Unit, users []int64) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Authorized on account: ", bot.Self.UserName)
	return &Bot{bot: bot, offset: offset, timeout: timeout, unit: unit, users: users}
}

func (b *Bot) CommandProcessor() error {
	u := tgbotapi.NewUpdate(b.offset)
	u.Timeout = b.timeout

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		reply := "dungeon bot do not know this command"
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			reply = "Hi Leatherman! Welcome to the club buddy!"
		case "getstate":
			if b.unit.State() == true {
				reply = "Gay bar is closed!"
			} else {
				reply = "Gay bar is open!"
			}
		case "seton":
			go func() {
				time.Sleep(time.Second * 60)
				b.unit.TurnOn()
			}()
			reply = "OK Leatherman! You have 60 seconds to make sure you have turned off the light, teapot and locked down the dungeon door!"
		case "setoff":
			b.unit.TurnOff()
			reply = "OK Leatherman! Enjoy the gay bar!"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		b.bot.Send(msg)
	}
	return nil
}

func (b *Bot) Notify(message models.Alarm) error {
	u := tgbotapi.NewUpdate(b.offset)
	u.Timeout = b.timeout

	log.Println(message.Message)

	for _, j := range b.users {
		msg := tgbotapi.NewMessage(j, message.Message)
		b.bot.Send(msg)
	}
	return nil
}

func (b *Bot) GWMessagesProcessor() {
	message := b.unit.GetMessage()
	if b.unit.State() == true {
		b.Notify(message)
	}
}

func (b *Bot) Ping() error {
	_, err := http.Get(b.unit.Url())
	if err != nil {
		return err
	}
	return nil
}
