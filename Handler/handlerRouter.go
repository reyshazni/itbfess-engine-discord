package Handler

import (
	"ITBFess/Handler/anon"
	"ITBFess/Handler/help"
	"github.com/bwmarrin/discordgo"
)

func HandlerRouter(dg *discordgo.Session) {
	dg.AddHandler(help.NewMemberHandler)
	dg.AddHandler(anon.AnonHandler)
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		//if strings.HasPrefix(m.Content, "/register") {
		//	registration.RegistrationHandler(s, m)
		//} else if strings.Contains(m.Content, "itb!") {
		//	FessHandler.FessHandler(s, m)
		//} else if strings.HasPrefix(m.Content, "/help") {
		//	help.HelpHandler(s, m)
		//}
	})
}
