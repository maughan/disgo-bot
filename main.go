package main

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

const token string = "NzU0MTAwMTEzNTAwMDEyNTY1.X1v0dQ.dIZvtwrl2vScDeSPQgjYS7CL27Q"
var BotID string

// func bumLogin () {
//     client := &http.Client{}
//     req, _ := http.NewRequest("POST", "https://bumble.com/mwebapi.phtml?SERVER_SUBMIT_PHONE_NUMBER", nil)
//     req.Header.Add("X-Desktop-web", "1")
//     req.Header.Add("Origin", "https://bumble.com")
//     req.Header.Add("Referer", "https://bumble.com/get-started")
//     req.Header.Add("Sec-Fetch-Mode", "cors")
//     req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
//     req.Header.Add("x-use-session-cookie", "1")
//     req.Body = ioutil.NopCloser(strings.NewReader(`{"version":1,"message_type":15,"message_id":12,"body":[{"message_type":15,"server_submit_phone_number":{"remember_me":true,"phone":"' + +447856916071 + '","password":"' + p + '","stats_data":""}}],"is_background":false}`))
//     resp, _ := client.Do(req)
//     defer resp.Body.Close()
//     if resp.StatusCode == http.StatusOK {
//         bodyBytes, err := ioutil.ReadAll(resp.Body)
//         if err != nil {
//             fmt.Println(err.Error())
//         }
//         bodyString := string(bodyBytes)
//         fmt.Println(bodyString)
//     }
// }

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