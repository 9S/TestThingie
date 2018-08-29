package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Handler for messageCreate
	dg.AddHandler(messageReceived)

	// Open the connection to Discord. At this stage, the bot will show up as online.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageReceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Don't run if the message was sent by the Bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(strings.ToLower(m.Content), Prefix+"Hi") {
		s.ChannelMessageSend(m.ChannelID, "Hello, world!")
	} else if strings.HasPrefix(strings.ToLower(m.Content), Prefix+"color") {
		Color(s, m)
	} else if strings.HasPrefix(strings.ToLower(m.Content), Prefix+"chucknorris") {
		ChuckNorris(s, m)
	} else if strings.HasPrefix(strings.ToLower(m.Content), Prefix+"die") {
		Die(s, m)
	}
}
