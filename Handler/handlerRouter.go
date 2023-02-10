package Handler

import (
	"ITBFess/Handler/FessHandler"
	"ITBFess/Handler/help"
	"ITBFess/Handler/registration"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HandlerRouter(dg *discordgo.Session) {
	dg.AddHandler(help.NewMemberHandler)
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, "/register") {
			registration.RegistrationHandler(s, m)
		} else if strings.Contains(m.Content, "itb!") {
			FessHandler.FessHandler(s, m)
		} else if strings.HasPrefix(m.Content, "/help") {
			help.HelpHandler(s, m)
		}
	})
}
