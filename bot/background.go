package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

// Starts up background tasks for the bot.
// It's important to not run this multiple times unless the running tasks somehow fail.
func (bot *Bot) LaunchTimedTasks() {
	go bot.ThreeAmEventTimer()
}

// Self-explanatory
func Is3AmSomewhere() (bool, Timezone) {
	fmt.Println("Is3AmSomewhere()")
	timestamp := time.Now().UTC()

	// minute should be 0, 15, 30, or 45
	isInCheckWindow := timestamp.Minute()%15 == 0
	if !isInCheckWindow {
		return false, Timezone{}
	}

	// figure out if any timezones are at 03:00 after accounting for offset
	for _, tz := range Timezones {
		offset := time.Hour*time.Duration(tz.OffsetHr) + time.Minute*time.Duration(tz.OffsetMin)
		localTime := timestamp.Add(offset)
		fmt.Printf("TIME FOR TIMEZONE %s: %s\n", tz.Name, localTime)
		if localTime.Hour() == 3 && localTime.Minute() == 0 {
			fmt.Printf("3AM IN TIMEZONE %s\n", strings.Replace(tz.Name, ".", ":", -1))
			return true, tz
		}
	}

	// default return value if nothing is found
	return false, Timezone{}
}

func (bot *Bot) ThreeAmEventTimer() {
	for {
		if is3am, tz := Is3AmSomewhere(); is3am {
			for _, guild := range bot.Session.State.Guilds {
				go bot.Dispatch3amEvent(tz, guild)
			}
			time.Sleep(time.Minute * 10)
		}
		time.Sleep(time.Second * 30)
	}
}

func (bot *Bot) Dispatch3amEvent(tz Timezone, guild *discordgo.Guild) {
	juan := "576168356823040010"
	revChannel := "577701223554351110"
	if guild.ID == juan {
		_, err := bot.Session.ChannelMessageSend(revChannel, "bong "+tz.Name)
		if err != nil {
			panic(err)
		}
	}
}
