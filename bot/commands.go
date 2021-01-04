package bot

import (
	"fmt"
	"github.com/nwunderly/disgo/commands"
	"github.com/nwunderly/oh-boy-3am/db"
)

func isAdmin(ctx *commands.Context) bool {
	if ctx.Author.ID == ctx.Guild.OwnerID {
		return true
	}
	for _, roleID := range ctx.Member.Roles {
		role, err := ctx.Session.State.Role(ctx.Guild.ID, roleID)
		if err != nil {
			continue
		}
		if role.Permissions & 8 != 0 {
			return true
		}
	}
	return false
}

func assertAdmin(ctx *commands.Context) bool {
	if isAdmin(ctx) {
		return true
	} else {
		_, _ = ctx.Send("You need admin permissions to do this!")
		return false
	}
}

func ViewConfig(ctx *commands.Context) error {
	channelID, ok := db.Database.GetChannelID(ctx.Guild.ID)
	if ok {
		_, _ = ctx.Send(fmt.Sprintf("<#%s>", channelID))
	} else {
		_, _ = ctx.Send("Could not find a configured channel for this guild.")
	}
	return nil
}

func SetConfig(ctx *commands.Context) error {
	if !assertAdmin(ctx) {
		return nil
	}
	if len(ctx.Args) != 1 {
		_, _ = ctx.Send("Please specify a channel.")
		return nil
	}
	channelID := ctx.Args[0]
	_, err := ctx.Session.State.Channel(channelID)
	if err != nil {
		_, _ = ctx.Send("Invalid channel ID.")
		return nil
	} else {
		_, _ = ctx.Send(fmt.Sprintf("Setting this guild's config channel to <#%s>.", channelID))
		err := db.Database.SetChannelID(ctx.Guild.ID, channelID)
		if err != nil {
			fmt.Println(err)
			_, _ = ctx.Send(err.Error())
		} else {
				_, _ = ctx.Send("Done.")
			}
		return nil
	}
}

func DelConfig(ctx *commands.Context) error {
	if !assertAdmin(ctx) {
		return nil
	}
	_, _ = ctx.Send("Removing this guild's configuration information.")
	err := db.Database.DelChannelID(ctx.Guild.ID)
	if err != nil {
		_, _ = ctx.Send(err.Error())
	} else {
		_, _ = ctx.Send("Done.")
	}
	return nil
}