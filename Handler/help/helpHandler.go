package help

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func HelpHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	msg := generateMsg()
	_, err := s.ChannelMessageSend(m.ChannelID, msg)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func NewMemberHandler(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	channel, err := s.UserChannelCreate(m.User.ID)
	if err != nil {
		fmt.Println("Error creating private channel: ", err)
		return
	}
	msg := generateMsg()
	_, _ = s.ChannelMessageSend(channel.ID, msg)
}

func generateMsg() string {
	msg := `**This bot will be terminated after 1st march 2023**
>>> **Help**
Greetings! We are excited to introduce ITBot/ITCord/ITBFess Discord. An ITBFess Alternative.

Please be advised that this platform is currently undergoing stress testing. While we strive to provide a seamless experience, it is possible that you may encounter bugs and face extended periods of downtime during this testing phase. We appreciate your understanding and patience as we work to improve the platform for a better user experience.

**Available Commands**
1. /register
To register, simply run the command with the designated prefix. 
This simple step will ensure that your account is eligible to participate and share your confessions.

2. itb! 
The command you are about to use is the classic ITBSFess/Menfess feature. It functions just like any regular confession command and serves as a platform for you to express yourself and share your thoughts with the community.

3. /help
The /help command is your ultimate guide to navigate through our platform. It provides you with a comprehensive list of all available commands and their respective functions, allowing you to make the most out of your time on our platform.

Whether you're new to our platform or simply need a refresher, the /help command is here to assist you. Simply type /help in any channel, and you will receive a detailed explanation of all the commands at your disposal.

**Future plans?**
[insert link]


p.s
We are not planning to take over Mio-chan's ITBFess. After 1st march 2023, all data will be deleted and all service will get terminated.
`
	return msg
}
