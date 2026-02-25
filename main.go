package main

import (
	"fmt"
	"log"
	"os/signal"
	"os"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "")
	if err != nil{
		log.Fatal(err)
	}
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
		if m.Author.ID == s.State.SessionID {
			return
		}

		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world")
		}
	})

	discord.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = discord.Open()
	if err != nil{
		log.Fatal(err)
	}
	defer discord.Close()

	fmt.Println("Bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}


