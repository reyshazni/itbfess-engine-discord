package Handler

import (
	"ITBFess/Handler/FessHandler"
	"github.com/bwmarrin/discordgo"
)

func HandlerRouter(dg *discordgo.Session) {
	dg.AddHandler(FessHandler.FessHandler)
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

	})
}
