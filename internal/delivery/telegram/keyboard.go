package telegram

import (
	"fmt"
	"testTGBot/internal/models"

	tgmodels "github.com/go-telegram/bot/models"
)

func BuildMainMenuKeyboard() *tgmodels.ReplyKeyboardMarkup {
	return &tgmodels.ReplyKeyboardMarkup{
		Keyboard: [][]tgmodels.KeyboardButton{
			{
				{Text: "🤖 Открыть ботов"},
			},
			{
				{Text: "ℹ️ Помощь"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}

func BuildCategoriesKeyboard(categories []models.Category) *tgmodels.InlineKeyboardMarkup {
	var rows [][]tgmodels.InlineKeyboardButton

	for _, cat := range categories {
		button := tgmodels.InlineKeyboardButton{
			Text:         cat.Icon + " " + cat.Name,
			CallbackData: "cat_" + fmt.Sprint(cat.ID),
		}
		rows = append(rows, []tgmodels.InlineKeyboardButton{button})
	}

	// Кнопка "Назад"
	rows = append(rows, []tgmodels.InlineKeyboardButton{
		{Text: "🔙 Назад", CallbackData: "back_main"},
	})

	return &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

// BuildBotsKeyboard создаёт Inline клавиатуру со списком ботов
func BuildBotsKeyboard(bots []models.Bot) *tgmodels.InlineKeyboardMarkup {
	var rows [][]tgmodels.InlineKeyboardButton

	for _, b := range bots {
		button := tgmodels.InlineKeyboardButton{
			Text:         b.Icon + " " + b.DisplayName,
			CallbackData: "bot_" + fmt.Sprint(b.ID),
		}
		rows = append(rows, []tgmodels.InlineKeyboardButton{button})
	}

	rows = append(rows, []tgmodels.InlineKeyboardButton{
		{Text: "🔙 Назад", CallbackData: "back_categories"},
	})

	return &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

func BuildBotDetailKeyboard(bot *models.Bot, categoryID int) *tgmodels.InlineKeyboardMarkup {
	return &tgmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgmodels.InlineKeyboardButton{
			{
				{Text: "🚀 Открыть бота", URL: bot.GetDeepLink()},
			},
			{
				{Text: "🔙 Назад к списку", CallbackData: "cat_" + fmt.Sprint(categoryID)},
			},
		},
	}
}
