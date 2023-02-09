package registration

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func RegistrationHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	botclient, err := s.User("@me")
	if m.GuildID == "" {
		return
	}
	if err != nil {
		fmt.Println("error getting bot user,", err)
		return
	}
	if m.Author.ID == botclient.ID {
		return
	}
}
