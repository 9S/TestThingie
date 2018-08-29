package main

// Prefix is the prefix used for bot commands.
const Prefix string = "$"

// LoadingForExternalAPIs shows loading message for external APIs.
//   Disadvantage: Messages show (edited).
const LoadingForExternalAPIs bool = true

// EmbedColor defines the color of the bar in front of embeds.
//   Can be one of the colors in EmbedColorList
const EmbedColor int = 0x3498db

// EmbedColorList includes all default discord colors with values.
// TODO: (Prio: LOW) Convert rest of list to hex values, to make it look nicer.
var EmbedColorList = map[string]int{
	"DEFAULT":     0x000000,
	"AQUA":        0x1ABC9C, //1752220,
	"GREEN":       0x2ECC71, //3066993,
	"BLUE":        0x3498DB, //3447003,
	"PURPLE":      0x9B59B6, //10181046,
	"GOLD":        0xF1C40F, //15844367, and so on.
	"ORANGE":      15105570,
	"RED":         15158332,
	"GREY":        9807270,
	"DARKER_GREY": 8359053,
	"NAVY":        3426654,
	"DARK_AQUA":   1146986,
	"DARK_GREEN":  2067276,
	"DARK_BLUE":   2123412,
	"DARK_PURPLE": 7419530,
	"DARK_GOLD":   12745742,
	"DARK_ORANGE": 11027200,
	"DARK_RED":    10038562,
	"DARK_GREY":   9936031,
	"LIGHT_GREY":  12370112,
	"DARK_NAVY":   2899536}
