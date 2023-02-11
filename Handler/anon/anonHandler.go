package anon

import (
	"ITBFess/Database"
	"ITBFess/Model/Entity"
	"ITBFess/Repository/anonRepository"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func AnonHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "/start") {
		startAnon(s, m)
	} else if strings.HasPrefix(m.Content, "/stop") {
		endAnon(s, m)
	} else if !m.Author.Bot {
		handleAnonInteraction(s, m)
	}
}

func startAnon(s *discordgo.Session, m *discordgo.MessageCreate) {
	db := Database.GetDatabase()
	senderID := m.ChannelID
	findMatch := Entity.Match{}
	err := db.Where("channel_id_first = ? or channel_id_second = ?", senderID, senderID).First(&findMatch).Error
	if err == nil {
		_, _ = s.ChannelMessageSend(senderID, "Kamu masih di percakapan\nKetik `/stop` untuk berhenti")
		return
	}
	_, err = s.ChannelMessageSend(senderID, "🕓 Sedang mencari calon pasangan...")
	if err != nil {
		return
	}
	receiverID, err := anonRepository.Get(anonRepository.All)
	if err != nil {
		err = anonRepository.Insert(anonRepository.All, senderID)
		return
	}
	match := Entity.Match{
		ChannelIdFirst:  m.ChannelID,
		ChannelIdSecond: receiverID,
	}
	err = db.Create(&match).Error
	if err != nil {
		log.Fatal(err)
	}
	_, _ = s.ChannelMessageSend(senderID, "💬 Pasangan ditemukan!")
	_, _ = s.ChannelMessageSend(receiverID, "💬 Pasangan ditemukan!")
}

func endAnon(s *discordgo.Session, m *discordgo.MessageCreate) {
	db := Database.GetDatabase()
	senderID := m.ChannelID
	findMatch := Entity.Match{}
	err := db.Where("channel_id_first = ? or channel_id_second = ?", senderID, senderID).First(&findMatch).Error
	if err != nil {
		_, err := anonRepository.Get(anonRepository.All)
		if err != nil {
			println(err.Error())
		}
		_, _ = s.ChannelMessageSend(senderID, "Kamu menghentikan percakapan\nKetik `/search` untuk mendapatkan partner baru!")
		return
	}
	_, _ = s.ChannelMessageSend(findMatch.ChannelIdSecond, "Percakapan dihentikan\nKetik `/search` untuk mendapatkan partner baru!")
	_, _ = s.ChannelMessageSend(findMatch.ChannelIdFirst, "Percakapan dihentikan\nKetik `/search` untuk mendapatkan partner baru!")
	if err = db.Delete(&findMatch).Error; err != nil {
		log.Fatal(err)
	}
}

func handleAnonInteraction(s *discordgo.Session, m *discordgo.MessageCreate) {
	senderID := m.ChannelID
	findMatch := Entity.Match{}
	db := Database.GetDatabase()
	err := db.Where("channel_id_first = ? or channel_id_second = ?", senderID, senderID).First(&findMatch).Error
	if err != nil {
		return
	}
	if findMatch.ChannelIdSecond == senderID {
		s.ChannelMessageSend(findMatch.ChannelIdFirst, m.Content)
		return
	}
	s.ChannelMessageSend(findMatch.ChannelIdSecond, m.Content)
}
