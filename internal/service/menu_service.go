package service

import (
	"fmt"
	"testTGBot/internal/models"
)

type MenuService struct {
	botService *BotService
}

func NewMenuService(botService *BotService) *MenuService {
	return &MenuService{botService: botService}
}

func (s *MenuService) FormatBotCard(bot *models.Bot) string {
	return fmt.Sprintf(
		"%s <b>%s</b>\n\n%s\n\n🔗 Нажми кнопку ниже, чтобы открыть бота",
		bot.Icon,
		bot.DisplayName,
		bot.Description,
	)
}

func (s *MenuService) FormatCategoryList(categories []models.Category) string {
	if len(categories) == 0 {
		return "Категории не найдены"
	}

	text := "📂 <b>Выбери категорию:</b>\n\n"
	for _, cat := range categories {
		text += fmt.Sprintf("%s %s\n", cat.Icon, cat.Name)
	}
	return text
}

func (s *MenuService) FormatBotsList(bots []models.Bot, categoryName string) string {
	if len(bots) == 0 {
		return "Боты не найдены"
	}

	text := fmt.Sprintf("🤖 <b>Боты в категории \"%s\":</b>\n\n", categoryName)
	text += "Выбери бота из списка ниже:"
	return text
}
func (s *MenuService) GetMainMenuText() string {
	return "🏠 <b>Главное меню</b>\n\n" +
		"Добро пожаловать в бот-навигатор!\n\n" +
		"Используй кнопки ниже для навигации:"
}
func (s *MenuService) GetHelpText() string {
	return "ℹ️ <b>Помощь</b>\n\n" +
		"<b>Как пользоваться ботом:</b>\n\n" +
		"1️⃣ Нажми \"🤖 Открыть ботов\" в главном меню\n" +
		"2️⃣ Выбери категорию ботов\n" +
		"3️⃣ Выбери нужного бота из списка\n" +
		"4️⃣ Нажми кнопку \"Открыть бота\"\n\n" +
		"<b>Доступные команды:</b>\n" +
		"/start - Главное меню\n" +
		"/help - Эта справка"
}
