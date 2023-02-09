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
			// Send a message to the Discord server
			fesses := []Entity.Menfess{}
			err := db.Find(&fesses).Error
			if err != nil {
				log.Fatal(err)
			}
			for _, fess := range fesses {
				msg := fmt.Sprintf("[AUTOMATED]\n>>> %s", fess.Message)
				_, err = dg.ChannelMessageSend(fess.AuthorChannelID, "Hai menfess kamu sudah di-tweet ya! 😃")
				_, err = dg.ChannelMessageSend(channel, msg)
				if err != nil {
					fmt.Println("Error sending message: ", err)
					return
				}
			}
			// Wait for the next interval
			db.Exec("TRUNCATE menfesses")
			time.Sleep(5 * time.Minute)
		}
	}()

}
