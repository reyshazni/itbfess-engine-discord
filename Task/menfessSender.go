package Task

import (
	"ITBFess/Database"
	"ITBFess/Model/Entity"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"time"
)

func MenfessSender(dg *discordgo.Session) {
	db := Database.GetDatabase()
	channel := os.Getenv("DC_FESS_CHANNEL")
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			println("Sending menfess..")
			fesses := []Entity.Menfess{}
			fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
			err := db.Where("created_at >= ?", fiveMinutesAgo).Find(&fesses).Error
			if err != nil {
				log.Fatal(err)
			}
			for _, fess := range fesses {
				msg := fmt.Sprintf("> [MENFESS]\n%s", fess.Message)
				author := Entity.User{}
				if err = db.First(&author, fess.AuthorID).Error; err != nil {
					log.Fatal(err)
				}
				author.DidSentFess = false
				if err = db.Save(&author).Error; err != nil {
					log.Fatal(err)
				}
				_, err = dg.ChannelMessageSend(fess.AuthorChannelID, "Hai menfess kamu sudah di-tweet ya! 😃")
				_, err = dg.ChannelMessageSend(channel, msg)
				if err != nil {
					fmt.Println("Error sending message: ", err)
				}
			}
		}
	}()

}
