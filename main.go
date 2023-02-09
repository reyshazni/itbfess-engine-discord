package main

import (
	"ITBFess/BotLoader"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	BotLoader.Loader()
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
