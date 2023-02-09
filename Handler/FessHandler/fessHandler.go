package FessHandler

import (
	"ITBFess/Database"
	"ITBFess/Model/Entity"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func FessHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.Contains(m.Content, "itb!") {
		botclient, err := s.User("@me")
		if m.GuildID != "" {
			return
		}
		if err != nil {
			fmt.Println("error getting bot user,", err)
			return
		}
		if m.Author.ID == botclient.ID {
			return
		}
		_, err = s.ChannelMessageSend(m.ChannelID, "Hai menfess kamu akan keluar sebentar lagi ya! 😃")
		db := Database.GetDatabase()
		menfess := Entity.Menfess{
			Message:         m.Content,
			AuthorID:        m.Author.ID,
			AuthorChannelID: m.ChannelID,
		}
		if db.Create(&menfess).Error != nil {
			log.Fatal("Cannot Add Menfess")
		}
	}
}
