package telegram

import (
	"context"
	"log/slog"
	"strconv"
	"strings"
	"testTGBot/internal/service"

	"github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
)

type Handler struct {
	botService  *service.BotService
	menuService *service.MenuService
	log         *slog.Logger
}

func NewHandler(botService *service.BotService, menuService *service.MenuService, log *slog.Logger) *Handler {
	return &Handler{
		botService:  botService,
		menuService: menuService,
		log:         log,
	}
}

func (h *Handler) HandleStart(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	text := h.menuService.GetMainMenuText()

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ParseMode:   tgmodels.ParseModeHTML,
		ReplyMarkup: BuildMainMenuKeyboard(),
	})

	if err != nil {
		h.log.Error("failed to send start message", "error", err)
	}
}

func (h *Handler) HandleHelp(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	text := h.menuService.GetHelpText()

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ParseMode:   tgmodels.ParseModeHTML,
		ReplyMarkup: BuildMainMenuKeyboard(),
	})

	if err != nil {
		h.log.Error("failed to send help message", "error", err)
	}
}

func (h *Handler) HandleOpenBots(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.Message == nil {
		return
	}

	chatID := update.Message.Chat.ID

	categories, err := h.botService.GetCategories(ctx)
	if err != nil {
		h.log.Error("failed to get categories", "error", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "❌ Ошибка при загрузке категорий",
		})
		return
	}

	text := h.menuService.FormatCategoryList(categories)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ParseMode:   tgmodels.ParseModeHTML,
		ReplyMarkup: BuildCategoriesKeyboard(categories),
	})

	if err != nil {
		h.log.Error("failed to send categories", "error", err)
	}
}
func (h *Handler) HandleCategoryCallback(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.CallbackQuery == nil {
		return
	}

	callback := update.CallbackQuery
	data := callback.Data
	parts := strings.Split(data, "_")
	if len(parts) != 2 {
		h.log.Error("invalid callback data format", "data", data)
		return
	}

	categoryID, err := strconv.Atoi(parts[1])
	if err != nil {
		h.log.Error("failed to parse category id", "error", err)
		return
	}

	bots, err := h.botService.GetBotsByCategory(ctx, categoryID)
	if err != nil {
		h.log.Error("failed to get bots by category", "error", err)
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: callback.ID,
			Text:            "Ошибка при загрузке ботов",
			ShowAlert:       true,
		})
		return
	}

	categories, _ := h.botService.GetCategories(ctx)
	var categoryName string
	for _, cat := range categories {
		if cat.ID == categoryID {
			categoryName = cat.Name
			break
		}
	}

	text := h.menuService.FormatBotsList(bots, categoryName)

	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      callback.Message.Message.Chat.ID,
		MessageID:   callback.Message.Message.ID,
		Text:        text,
		ParseMode:   tgmodels.ParseModeHTML,
		ReplyMarkup: BuildBotsKeyboard(bots),
	})

	if err != nil {
		h.log.Error("failed to edit message", "error", err)
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
	})
}

func (h *Handler) HandleBotCallback(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.CallbackQuery == nil {
		return
	}

	callback := update.CallbackQuery
	data := callback.Data

	parts := strings.Split(data, "_")
	if len(parts) != 2 {
		h.log.Error("invalid callback data format", "data", data)
		return
	}

	botID, err := strconv.Atoi(parts[1])
	if err != nil {
		h.log.Error("failed to parse bot id", "error", err)
		return
	}

	botInfo, err := h.botService.GetBotByID(ctx, botID)
	if err != nil {
		h.log.Error("failed to get bot by id", "error", err)
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: callback.ID,
			Text:            "Ошибка при загрузке информации о боте",
			ShowAlert:       true,
		})
		return
	}

	text := h.menuService.FormatBotCard(botInfo)

	// Редактируем сообщение
	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      callback.Message.Message.Chat.ID,
		MessageID:   callback.Message.Message.ID,
		Text:        text,
		ParseMode:   tgmodels.ParseModeHTML,
		ReplyMarkup: BuildBotDetailKeyboard(botInfo, botInfo.CategoryID),
	})

	if err != nil {
		h.log.Error("failed to edit message", "error", err)
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
	})
}

func (h *Handler) HandleBackToCategories(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.CallbackQuery == nil {
		return
	}

	callback := update.CallbackQuery

	categories, err := h.botService.GetCategories(ctx)
	if err != nil {
		h.log.Error("failed to get categories", "error", err)
		return
	}

	text := h.menuService.FormatCategoryList(categories)

	_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      callback.Message.Message.Chat.ID,
		MessageID:   callback.Message.Message.ID,
		Text:        text,
		ParseMode:   tgmodels.ParseModeHTML,
		ReplyMarkup: BuildCategoriesKeyboard(categories),
	})

	if err != nil {
		h.log.Error("failed to edit message", "error", err)
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
	})
}

func (h *Handler) HandleBackToMain(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	if update.CallbackQuery == nil {
		return
	}

	callback := update.CallbackQuery
	text := h.menuService.GetMainMenuText()

	_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    callback.Message.Message.Chat.ID,
		MessageID: callback.Message.Message.ID,
		Text:      text,
		ParseMode: tgmodels.ParseModeMarkdown,
	})

	if err != nil {
		h.log.Error("failed to edit message", "error", err)
	}

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callback.ID,
	})
}
