package registration

import (
	"ITBFess/Database"
	"ITBFess/Model/Entity"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func RegistrationHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	botclient, err := s.User("@me")
	if m.Author.Bot {
		return
	}
	if len(m.GuildID) > 0 {
		_, err = s.ChannelMessageSend(m.ChannelID, "You are not allowed to perform this function on a server channel")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if err != nil {
		fmt.Println("error getting bot user,", err)
		return
	}
	if m.Author.ID == botclient.ID {
		fmt.Println("author id kok sama?")
		return
	}
	author := Entity.User{
		AuthorID:    m.Author.ID,
		IsVerified:  true,
		IsAdmin:     false,
		DidSentFess: false,
	}
	db := Database.GetDatabase()
	if err := db.Create(&author).Error; err != nil {
		_, err = s.ChannelMessageSend(m.ChannelID, "Kamu sudah terdaftar!")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	_, err = s.ChannelMessageSend(m.ChannelID, "Selemat! Kamu berhasil ter-register ke sistem")
	if err != nil {
		log.Fatal(err)
	}
	return
}
