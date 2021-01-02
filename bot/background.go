package bot

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

// Starts up background tasks for the bot.
// It's important to not run this multiple times unless the running tasks somehow fail.
func (bot *Bot) LaunchTimedTasks() {
	go bot.ThreeAmEventTimer()
}

// Self-explanatory
func Is3AmSomewhere() (bool, Timezone) {
	//Todo: this
	return false, Timezone{}
}

func (bot *Bot) ThreeAmEventTimer() {
	for {
		if is3am, tz := Is3AmSomewhere(); is3am {
			for _, guild := range bot.Session.State.Guilds {
				go bot.Dispatch3amEvent(tz, guild)
			}
			time.Sleep(time.Minute*10)
		}
		time.Sleep(time.Second*30)
	}
}

func (bot *Bot) Dispatch3amEvent(tz Timezone, guild *discordgo.Guild) {
	//Todo: this
}