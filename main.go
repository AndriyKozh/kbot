package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "kbot",
		Short: "Telegram bot written in Go",
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start Telegram bot",

		RunE: func(cmd *cobra.Command, args []string) error {
			token := os.Getenv("TELE_TOKEN")

			if token == "" {
				return fmt.Errorf("TELE_TOKEN is not set")
			}

			bot, err := telebot.NewBot(telebot.Settings{
				Token: token,
				Poller: &telebot.LongPoller{
					Timeout: 10 * time.Second,
				},
			})

			if err != nil {
				return fmt.Errorf("failed to create bot: %w", err)
			}

			bot.Handle("/start", func(c telebot.Context) error {
				return c.Send("Привіт! Я Telegram-бот на Go!")
			})

			bot.Handle("/help", func(c telebot.Context) error {
				return c.Send("Команди: /start, /help. Напиши hello!")
			})

			bot.Handle(telebot.OnText, func(c telebot.Context) error {
				text := strings.ToLower(strings.TrimSpace(c.Text()))

				switch text {
				case "hello", "привіт":
					return c.Send("Вітаю! Бот працює!")
				default:
					return c.Send("Отримав повідомлення: " + c.Text())
				}
			})

			fmt.Println("Telegram bot is running...")
			bot.Start()

			return nil
		},
	}

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
