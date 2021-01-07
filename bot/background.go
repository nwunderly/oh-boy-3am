package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nwunderly/oh-boy-3am/db"
	"time"
)

// Starts up background tasks for the bot.
// It's important to not run this multiple times unless the running tasks somehow fail.
func (bot *Bot) LaunchTimedTasks() {
	go bot.ThreeAmEventTimer()
}

// Self-explanatory
func Is3AmSomewhere() (bool, Timezone) {
	timestamp := time.Now().UTC()
	// minute should be 0, 15, 30, or 45
	isInCheckWindow := timestamp.Minute()%15 == 0
	if !isInCheckWindow {
		return false, Timezone{}
	}

	// figure out if any timezones are at 03:00 after accounting for offset
	for _, tz := range Timezones {
		if tz.Is3Am() {
			return true, tz
		}
	}

	// default return value if nothing is found
	return false, Timezone{}
}

func (bot *Bot) ThreeAmEventTimer() {
	for {
		if is3am, tz := Is3AmSomewhere(); is3am {
			fmt.Println("3am detected in", tz)
			for _, guild := range bot.Session.State.Guilds {
				fmt.Println("\tdispatching 3am event to guild", guild.ID)
				bot.Dispatch3amEvent(tz, guild)
			}
			time.Sleep(time.Minute * 10)
		}
		time.Sleep(time.Second * 30)
	}
}

func (bot *Bot) Dispatch3amEvent(tz Timezone, guild *discordgo.Guild) {
	channelID, ok := db.Database.GetChannelID(guild.ID)
	if ok {
		fmt.Println("\t\tguild", guild.ID, "has 3am events configured for channel", channelID)
		_, err := bot.Session.ChannelMessageSend(channelID, fmt.Sprintf("OH BOY 3AM (%s)", tz.Name))
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("\t\tcould not find a channel for guild", guild.ID)
	}
}
