
# Kbot — Telegram Bot on Golang

Telegram-бот, розроблений на Go в рамках практичного завдання курсу DevOps & Kubernetes.

## Technologies

- Golang
- Cobra CLI
- Telebot v4
- Telegram Bot API
- GitHub Codespaces

## Features

- `/start` — привітання користувача
- `/help` — список доступних команд
- `hello` або `привіт` — відповідь на привітання
- Інші текстові повідомлення — бот повторює отриманий текст

## Installation

Клонування репозиторію:

```bash
git clone https://github.com/AndriyKozh/kbot.git
cd kbot
```

Завантаження залежностей:

```bash
go mod download
```

Компіляція:

```bash
go build -o kbot .
```

## Configuration

Створіть Telegram-бота через [BotFather](https://t.me/BotFather) та отримайте API Token.

Встановіть змінну середовища в Bash:

```bash
read -rs -p "Telegram token: " TELE_TOKEN
echo
export TELE_TOKEN
```

Не додавайте токен до репозиторію.

## Usage

Запуск бота:

```bash
./kbot start
```

Або без попередньої компіляції:

```bash
go run . start
```

Перегляд CLI-команд:

```bash
./kbot --help
```

## Telegram Bot

https://t.me/andriy_devops_kbot

## Author

GitHub: [AndriyKozh](https://github.com/AndriyKozh)
