package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Color writes out a embeded message to a discord
func Color(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Extract Color from message
	text := m.Content
	color := strings.Trim(text, " ")
	// Use last 6 digits. Solves Problem of people writing
	//   #ffffff, 0xffffff or only ffffff if they want white.
	color = color[len(color)-6:]
	colorInt, err := strconv.ParseInt(color, 16, 64)

	if err != nil {
		HelpColor(s, m.ChannelID)
		return
	}

	message := discordgo.MessageEmbed{
		Title:       "Hey! Here's your color!",
		Description: "And a fancy one it is, color: #" + string(color),
		Timestamp:   time.Now().Format("2006-01-02 15:04:05"),
		Color:       int(colorInt)}
	s.ChannelMessageSendEmbed(m.ChannelID, &message)
}

// ChuckNorris facts for everyone
func ChuckNorris(s *discordgo.Session, m *discordgo.MessageCreate) {

	msg, _ := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Loading...",
		Description: "from api.chucknorris.io",
		Timestamp:   time.Now().UTC().Format("2006-01-02 15:04:05"),
		Color:       EmbedColor})

	json, err := GetJSONFromURL("https://api.chucknorris.io/jokes/random")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "`Oops. Something went wrong when contacting chucknorris.io. Try again later`")
	}

	joke, err := GetKeyFromJSON(json, "value", false)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "`Oops. Something went wrong with processing the response.`")
	}

	message := discordgo.MessageEmbed{
		Title:       "Fun-Fact about Chuck Norris:",
		Description: joke,
		Timestamp:   time.Now().UTC().Format("2006-01-02 15:04:05"),
		Color:       EmbedColor}
	s.ChannelMessageEditEmbed(m.ChannelID, msg.ID, &message)

}

// Die kills the Bot.
func Die(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "PANIC! I'm DYING!")
	panic("Dead.")
}
