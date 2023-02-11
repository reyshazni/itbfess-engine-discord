package BotLoader

import (
	"ITBFess/Database/redisdb"
	"ITBFess/Handler"
	"ITBFess/Task"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

func Loader() {
	var token string = os.Getenv("DC_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	redisdb.GetClient()
	if err != nil {
		log.Fatal("Problem loading discord bot")
		return
	}
	dg.Identify.Intents |= discordgo.IntentGuildMembers
	dg.Identify.Intents |= discordgo.IntentsGuildMessages
	Handler.HandlerRouter(dg)
	Task.MenfessSender(dg)

	if err = dg.Open(); err != nil {
		log.Fatal(err)
		return
	}
}
