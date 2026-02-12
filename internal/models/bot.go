package models

import "time"

// Bot представляет информацию о Telegram боте
type Bot struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	Icon        string    `json:"icon"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Bot) GetDeepLink() string {
	return "https://t.me/" + b.Username
}
