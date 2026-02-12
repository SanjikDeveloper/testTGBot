package service

import (
	"context"
	"fmt"
	"testTGBot/internal/models"
	"testTGBot/internal/repository"
)

type BotService struct {
	repo repository.BotRepository
}

func NewBotService(repo repository.BotRepository) *BotService {
	return &BotService{repo: repo}
}

func (s *BotService) GetAllBots(ctx context.Context) ([]models.Bot, error) {
	bots, err := s.repo.GetAllBots(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get all bots: %w", err)
	}
	return bots, nil
}

func (s *BotService) GetBotsByCategory(ctx context.Context, categoryID int) ([]models.Bot, error) {
	bots, err := s.repo.GetBotsByCategory(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get bots by category: %w", err)
	}
	return bots, nil
}

func (s *BotService) GetCategories(ctx context.Context) ([]models.Category, error) {
	categories, err := s.repo.GetCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get categories: %w", err)
	}
	return categories, nil
}

func (s *BotService) GetBotByID(ctx context.Context, id int) (*models.Bot, error) {
	bot, err := s.repo.GetBotByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get bot by id: %w", err)
	}
	return bot, nil
}
