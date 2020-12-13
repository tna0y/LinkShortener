package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"service/pkg/actions"
	"service/pkg/entities"
)

const addCommandValue = "add"

func (b *Bot) addCommand(ctx context.Context, args string, chatID int64) (err error) {

	requester := entities.NewRequester(strconv.FormatInt(chatID, 10))

	var callArgs actions.CreateLinkArgs
	callArgs, err = b.parseAddArgs(args)
	if err != nil {
		return
	}

	var result entities.Link
	result, err = b.actions.CreateLink(ctx, callArgs, requester)
	if err != nil {
		return
	}

	reply := tgbotapi.NewMessage(chatID, b.buildAddResponse(result))
	_, err = b.botAPI.Send(reply)
	return
}

func (b *Bot) parseAddArgs(args string) (result actions.CreateLinkArgs, err error) {
	argsSplit := strings.Split(args, " ")
	if len(argsSplit) > 3 || len(argsSplit) < 2 || args == "" {
		err = errInvalidArguments
		return
	}

	linkID := argsSplit[0]
	target := argsSplit[1]

	var ttl int64
	if len(argsSplit) >= 3 {
		ttl, err = strconv.ParseInt(argsSplit[2], 10, 64)
		if err != nil {
			err = errInvalidArguments
			return
		}
	}

	return actions.CreateLinkArgs{
		ShortID: linkID,
		Target:  target,
		TTL:     ttl,
	}, nil
}

func (b *Bot) buildAddResponse(result entities.Link) string {
	url := fmt.Sprintf("%s/%s", b.config.BaseURL, result.ShortID)
	return fmt.Sprintf(linkCreatedMessage, url)
}
