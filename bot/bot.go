package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nwunderly/disgo/commands"
)

type Bot struct {
	*commands.Bot
	Debug bool
}

func New(prefix string, token string, debug bool) *Bot {
	cmdBot, err := commands.NewBot(prefix, token)
	if err != nil {
		panic(err)
	}

	bot := &Bot{
		Bot: cmdBot,
		Debug: debug,
	}

	return bot
}

func (bot *Bot) Run() {
	// Channel to send ready signal to
	whenReady := make(chan bool)

	// Setup trigger to launch tasks that are on timers
	bot.Session.AddHandlerOnce(
		func(session *discordgo.Session, ready *discordgo.Ready) {
			whenReady <- true
			close(whenReady)
		})

	// Start timed tasks when the bot is ready
	go func() {
		<-whenReady
		fmt.Println("Bot is ready. Launching timed tasks.")
		bot.LaunchTimedTasks()
	}()

	bot.Bot.Run()
}
