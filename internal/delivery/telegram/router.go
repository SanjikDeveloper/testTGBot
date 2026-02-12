package telegram

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) Route(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {

	if update.CallbackQuery != nil {
		r.handleCallback(ctx, b, update)
		return
	}

	if update.Message != nil {
		r.handleMessage(ctx, b, update)
		return
	}
}

func (r *Router) handleCallback(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	data := update.CallbackQuery.Data

	switch {
	case strings.HasPrefix(data, "cat_"):
		r.handler.HandleCategoryCallback(ctx, b, update)
	case strings.HasPrefix(data, "bot_"):
		r.handler.HandleBotCallback(ctx, b, update)
	case data == "back_categories":
		r.handler.HandleBackToCategories(ctx, b, update)
	case data == "back_main":
		r.handler.HandleBackToMain(ctx, b, update)
	}
}

func (r *Router) handleMessage(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	text := update.Message.Text

	if strings.HasPrefix(text, "/") {
		switch text {
		case "/start":
			r.handler.HandleStart(ctx, b, update)
		case "/help":
			r.handler.HandleHelp(ctx, b, update)
		}
		return
	}

	switch text {
	case "🤖 Открыть ботов":
		r.handler.HandleOpenBots(ctx, b, update)
	case "ℹ️ Помощь":
		r.handler.HandleHelp(ctx, b, update)
	default:
		// Если неизвестная команда, показываем главное меню
		r.handler.HandleStart(ctx, b, update)
	}
}
