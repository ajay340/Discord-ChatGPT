package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajay340/Discord-ChatGPT/discord"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	botSession         *discordgo.Session
	registeredCommands []*discordgo.ApplicationCommand
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	var bot_token string = os.Getenv("BOT_TOKEN")
	botSession, err = discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
		return
	}
}

func addCommands(session *discordgo.Session) {
	log.Println("Adding commands...")
	registeredCommands = make([]*discordgo.ApplicationCommand, len(discord.COMMANDS))
	for i, command := range discord.COMMANDS {
		cmd, err := botSession.ApplicationCommandCreate(botSession.State.User.ID, "", command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", command.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

func removeCommands(session *discordgo.Session) {
	for _, command := range registeredCommands {
		err := session.ApplicationCommandDelete(session.State.User.ID, "", command.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", command.Name, err)
		}
	}
}

func main() {

	botSession.AddHandler(discord.CommandInteractions)

	err := botSession.Open()
	if err != nil {
		log.Fatalln("error opening connection,", err)
	}

	addCommands(botSession)

	botSession.Identify.Intents = discordgo.IntentsGuildMessages

	defer botSession.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	removeCommands(botSession)
}
