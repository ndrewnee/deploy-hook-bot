# 🤖 deploy-hook-bot

[![Go](https://github.com/ndrewnee/deploy-hook-bot/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ndrewnee/deploy-hook-bot/actions/workflows/go.yml)

[Telegram bot](https://t.me/lesswrong_bot) for sending notification when something is deployed on Heroku.

## 🧑‍💻 Run locally

Register new bot at https://t.me/BotFather or use previously created one.

Take bot access token.

Before running copy sample file and replace env vars with your credentials

```sh
cp .env.sample .env
source .env
```

Run application locally:

```sh
make run
```

Also you can run bot with redis in docker compose:

```sh
docker-compose up
```

## 👷 Build

Build binary

```sh
make build
```

## 🧪 Testing

Run unit tests

```sh
make test
```

Run integration tests

```sh
make test_integration
```

## 🖍 Lint

Run linters

```sh
make lint
```

## 🛥 Deployment

Automatic CI/CD pipelines are building and testing the bot on each PR.

Demo bot is deployed to production on Heroku on merge to master.

To deploy your app on Heroku read [documentation](https://devcenter.heroku.com/articles/getting-started-with-go?singlepage=true).

```sh
brew install heroku/brew/heroku

heroku login
heroku create deploy-hook-bot
heroku config:set TOKEN=<token>
heroku config:set TELEGRAM_CHAT_ID=<chat_id>
heroku webhooks:add -i api:build -l notify -u https://deploy-hook-bot.herokuapp.com/hooks # To add deploy hook

git push heroku main
```

If application is already setup just run:

```sh
make deploy
```

## 🛠 Environment variables

| Env var          | Type    | Description               | Default |
| ---------------- | ------- | ------------------------- | ------- |
| PORT             | String  | Port for server           | 9998    |
| TOKEN            | String  | Telegram bot access token |         |
| TELEGRAM_CHAT_ID | Integer | Telegram chat id          |         |
| DEBUG            | Boolean | Enable debug mode         | false   |
