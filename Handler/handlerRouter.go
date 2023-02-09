package Handler

import (
	"ITBFess/Handler/FessHandler"
	"github.com/bwmarrin/discordgo"
)

func HandlerRouter(dg *discordgo.Session) {
	dg.AddHandler(FessHandler.FessHandler)
}
