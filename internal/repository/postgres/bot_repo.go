package postgres

import (
	"context"
	"fmt"
	"testTGBot/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BotRepo struct {
	pool *pgxpool.Pool
}

func NewBotRepo(pool *pgxpool.Pool) *BotRepo {
	return &BotRepo{pool: pool}
}

func (r *BotRepo) GetAllBots(ctx context.Context) ([]models.Bot, error) {
	query := `
		SELECT id, username, display_name, description, category_id, icon, is_active, created_at, updated_at
		FROM bots
		WHERE is_active = true
		ORDER BY display_name
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query bots: %w", err)
	}
	defer rows.Close()

	var bots []models.Bot
	for rows.Next() {
		var bot models.Bot
		err := rows.Scan(
			&bot.ID,
			&bot.Username,
			&bot.DisplayName,
			&bot.Description,
			&bot.CategoryID,
			&bot.Icon,
			&bot.IsActive,
			&bot.CreatedAt,
			&bot.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bot: %w", err)
		}
		bots = append(bots, bot)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return bots, nil
}

func (r *BotRepo) GetBotsByCategory(ctx context.Context, categoryID int) ([]models.Bot, error) {
	query := `
		SELECT id, username, display_name, description, category_id, icon, is_active, created_at, updated_at
		FROM bots
		WHERE category_id = $1 AND is_active = true
		ORDER BY display_name
	`

	rows, err := r.pool.Query(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to query bots by category: %w", err)
	}
	defer rows.Close()

	var bots []models.Bot
	for rows.Next() {
		var bot models.Bot
		err := rows.Scan(
			&bot.ID,
			&bot.Username,
			&bot.DisplayName,
			&bot.Description,
			&bot.CategoryID,
			&bot.Icon,
			&bot.IsActive,
			&bot.CreatedAt,
			&bot.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan bot: %w", err)
		}
		bots = append(bots, bot)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return bots, nil
}

func (r *BotRepo) GetCategories(ctx context.Context) ([]models.Category, error) {
	query := `
		SELECT id, name, description, icon
		FROM categories
		ORDER BY name
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name, &cat.Description, &cat.Icon)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, cat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return categories, nil
}

func (r *BotRepo) GetBotByID(ctx context.Context, id int) (*models.Bot, error) {
	query := `
		SELECT id, username, display_name, description, category_id, icon, is_active, created_at, updated_at
		FROM bots
		WHERE id = $1
	`

	var bot models.Bot
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&bot.ID,
		&bot.Username,
		&bot.DisplayName,
		&bot.Description,
		&bot.CategoryID,
		&bot.Icon,
		&bot.IsActive,
		&bot.CreatedAt,
		&bot.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get bot by id: %w", err)
	}

	return &bot, nil
}
