package BotLoader

import (
	"ITBFess/Handler"
	"ITBFess/Task"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

func Loader() {
	var token string = os.Getenv("DC_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Problem loading discord bot")
		return
	}
	if err = dg.Open(); err != nil {
		log.Fatal(err)
		return
	}
	Handler.HandlerRouter(dg)
	Task.MenfessSender(dg)
}
