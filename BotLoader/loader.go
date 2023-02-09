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
	if dg.Open() != nil {
		log.Fatal("Cannot Open DC Bot!")
		return
	}
	Handler.HandlerRouter(dg)
	Task.MenfessSender(dg)
}
