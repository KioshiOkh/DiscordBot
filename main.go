package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Token  string
	Prefix string

	config *configStruct
)

type configStruct struct {
	Token  string `json:"Token"`
	Prefix string `json:"Prefix"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading config.json...")
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	Token = config.Token
	Prefix = config.Prefix

	return nil
}

var BotID string
var goBot *discordgo.Session

func Start() {
	ReadConfig()

	goBot, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err)
	}

	u, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == Prefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

func main() {

	Start()
	<-make(chan struct{})
	return
}
