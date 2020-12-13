package bot

import (
	"context"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"service/pkg/actions"
	"service/pkg/config"
)

type Bot struct {
	actions *actions.Actions
	botAPI  *tgbotapi.BotAPI
	config  config.Config
}

func NewBot(cfg config.Config, act *actions.Actions) *Bot {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	return &Bot{
		botAPI:  bot,
		config:  cfg,
		actions: act,
	}
}

func (b *Bot) Name() string {
	return serviceName
}

func (b *Bot) Run(ctx context.Context) error {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.botAPI.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for {
		select {
		case update := <-updates:
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}
			err = b.route(ctx, update.Message)
			if err != nil {
				log.Printf("error handling message: %s", err.Error())
				_ = b.error(update.Message.Chat.ID, err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (b *Bot) route(ctx context.Context, msg *tgbotapi.Message) (err error) {
	if !msg.IsCommand() {
		return b.help(msg.Chat.ID)
	}

	switch msg.Command() {
	case addCommandValue:
		return b.addCommand(ctx, msg.CommandArguments(), msg.Chat.ID)
	case listCommandValue:
		return b.listCommand(ctx, msg.CommandArguments(), msg.Chat.ID)
	case removeCommandValue:
		return b.removeCommand(ctx, msg.CommandArguments(), msg.Chat.ID)
	default:
		return b.help(msg.Chat.ID)
	}
}

func (b *Bot) help(chatID int64) (err error) {
	reply := tgbotapi.NewMessage(chatID, helpMessage)
	_, err = b.botAPI.Send(reply)
	return err
}

func (b *Bot) error(chatID int64, errIn error) (err error) {
	reply := tgbotapi.NewMessage(chatID, errIn.Error())
	_, err = b.botAPI.Send(reply)
	return err
}
