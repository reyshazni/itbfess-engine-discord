package Handler

import (
	"ITBFess/Handler/FessHandler"
	"ITBFess/Handler/registration"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func HandlerRouter(dg *discordgo.Session) {
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, "/register") {
			registration.RegistrationHandler(s, m)
		}
		if strings.Contains(m.Content, "itb!") {
			FessHandler.FessHandler(s, m)
		}
	})
}
