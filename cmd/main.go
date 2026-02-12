package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"testTGBot/internal/delivery/telegram"
	"testTGBot/internal/repository/postgres"
	"testTGBot/internal/service"
	"testTGBot/pkg/config"
	"testTGBot/pkg/logger"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		slog.Error("error loading env variables: %s", err.Error())
		return
	}
	cfg := &config.Config{}
	if err := config.ReadEnvConfig(cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	logger := logger.New(cfg.Env)
	logger.Info("starting telegram bot navigator", "env", cfg.Env)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	pool, err := postgres.New(dsn)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		log.Fatalf("database connection failed: %v", err)
	}
	defer pool.Close()
	logger.Info("connected to database successfully")

	botRepo := postgres.NewBotRepo(pool)

	botService := service.NewBotService(botRepo)
	menuService := service.NewMenuService(botService)

	handler := telegram.NewHandler(botService, menuService, logger)
	router := telegram.NewRouter(handler)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(router.Route),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		logger.Error("failed to create bot", "error", err)
		log.Fatalf("bot creation failed: %v", err)
	}

	logger.Info("bot started successfully")
	b.Start(ctx)

	logger.Info("bot stopped gracefully")
}

//
