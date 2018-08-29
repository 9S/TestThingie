package main

import (
	"github.com/bwmarrin/discordgo"
)

// HelpColor proves help for the command Color
func HelpColor(s *discordgo.Session, channelID string) {
	helpText := Prefix + "colors #ffffff\nSpecify a color and see what it looks like."
	s.ChannelMessageSend(channelID, helpText)
}
