package main

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

const token string = "NzU0MTAwMTEzNTAwMDEyNTY1.X1v0dQ.dIZvtwrl2vScDeSPQgjYS7CL27Q"
var BotID string

func main () {
    dg, err := discordgo.New("Bot " + token)

    if err != nil {
        fmt.Println(err.Error(), "first")
        return
    }

    u, err := dg.User("@me")
    if err != nil {
        fmt.Println(err.Error())
    }

    BotID = u.ID

    dg.AddHandler(messageHandler)

    err = dg.Open()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Bot is running")


    <-make(chan struct{})
    return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
    var isOnAdventure bool

    if m.Author.ID == BotID {
        return
    }

    if m.Content == "ping" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "pong")
    }

    if m.Content == "!help" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "Hi! Here is the stuff I can do: \n\n `ping`: play a delightful game of table tennis \n `!adventure`: Begin an adventure")
    }

    if m.Content == "!adventure" {
        _, _ = s.ChannelMessageSend(m.ChannelID, "So you'd like to start an adventure? Let's see where we were...")
        isOnAdventure = true
        _, _ = s.ChannelMessageSend(m.ChannelID, "Adventure Controls: \n\n `!character`:  See your character sheet, will start character creation if you have no characters.\n `!begin`: Head off on your adventure. \n `!map`: Displays a map if possible.")


    }
    if m.Content == "!status" {
                    if isOnAdventure == true {
                        _,_ = s.ChannelMessageSend(m.ChannelID, "You are on an adventure!")
                    } else {
                        _,_ = s.ChannelMessageSend(m.ChannelID, "You are not on an adventure!")
                    }
                }

                if m.Content == "!character" {
                    _, _ = s.ChannelMessageSend(m.ChannelID, "Let's create your character sheet!")
                }

}
