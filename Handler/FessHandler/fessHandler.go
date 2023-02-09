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
		db := Database.GetDatabase()
		user := Entity.User{}
		err = db.First(&user, m.Author.ID).Error
		if err != nil {
			_, err = s.ChannelMessageSend(m.ChannelID, "Ngirim menfess tapi belum verified, cuih!")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		if user.DidSentFess {
			_, err = s.ChannelMessageSend(m.ChannelID, "Ih! kamu jangan spam dong 🥺, tunggu menfess kamu sebelumnya terkirim dulu oke?")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		_, err = s.ChannelMessageSend(m.ChannelID, "Hai menfess kamu akan keluar sebentar lagi ya! 😃")
		if err != nil {
			log.Fatal(err)
		}
		user.DidSentFess = true
		if err = db.Save(&user).Error; err != nil {
			log.Fatal(err)
		}
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
