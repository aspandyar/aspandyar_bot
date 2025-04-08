# aspandyar_bot

A Telegram bot designed to provide various functionalities through available commands.

## Features

- Easy-to-use Telegram bot with customizable commands.
- Dockerized setup for seamless deployment.
- Scalable and maintainable architecture.

## Prerequisites

- Docker installed on your system.
- A valid Telegram bot token. You can obtain one by creating a bot through [BotFather](https://core.telegram.org/bots#botfather).
- Chat GPT token
- app.env file, completed, like in dist.env

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/aspandyar_bot.git
    cd aspandyar_bot
    ```

2. Create a `.env` file in the project root and add your Telegram bot token:

3. Start the bot using Docker Compose:

    ```bash
    docker compose up
    ```

    Or just run:

    ```bash
    go run main.go
    ```

4. Your bot is now running and ready to use!

## Available Commands

- `/start` - Initialize the bot and get a welcome message.
- `/help` - List all available commands and their descriptions.
- `/begin` - Start chat gpt speaking option.

## Contributing

Contributions are welcome! Feel free to fork the repository, make changes, and submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
