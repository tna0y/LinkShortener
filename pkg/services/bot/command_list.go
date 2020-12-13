package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"service/pkg/entities"
)

const listCommandValue = "list"

func (b *Bot) listCommand(ctx context.Context, args string, chatID int64) (err error) {
	requester := entities.NewRequester(strconv.FormatInt(chatID, 10))

	var result []entities.Link
	result, err = b.actions.ListLinks(ctx, requester)
	if err != nil {
		return
	}

	reply := tgbotapi.NewMessage(chatID, b.buildListResponse(result))
	_, err = b.botAPI.Send(reply)
	return
}

func (b *Bot) buildListResponse(result []entities.Link) string {
	var rows []string
	for _, link := range result {
		var expires string
		if !link.Expires.IsZero() {
			expires = fmt.Sprintf(linkListExpiresMessage, link.Expires.String())
		}

		url := fmt.Sprintf("%s/%s", b.config.BaseURL, link.ShortID)

		rows = append(rows, fmt.Sprintf(linkListItemMessage, url, link.Target, expires))
	}

	return fmt.Sprintf(linkListMessage, len(result), strings.Join(rows, ""))
}
