package repository

import (
	"context"
	"testTGBot/internal/models"
)

type BotRepository interface {
	GetAllBots(ctx context.Context) ([]models.Bot, error)

	GetBotsByCategory(ctx context.Context, categoryID int) ([]models.Bot, error)

	GetCategories(ctx context.Context) ([]models.Category, error)

	GetBotByID(ctx context.Context, id int) (*models.Bot, error)
}
