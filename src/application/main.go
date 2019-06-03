package main

import (
    "gopkg.in/telegram-bot-api.v4"
    "application/config"
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "os"
)

var c *config.Config

func init() {
    conf := flag.String("conf", "conf/conf.json", "a string")

    flag.Parse()

    log.Printf("read conf => %s", *conf)

    file, err := os.Open(*conf)
    if err != nil {
        log.Println(err)
    }
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&c)
    if err != nil {
        log.Println(err)
    }
}

func main() {
    bot, err := tgbotapi.NewBotAPI(fmt.Sprintf("%s:%s", c.Telegram.Bot, c.Telegram.Token))
    if err != nil {
        log.Panic(err)
    }

    bot.Debug = false

    log.Printf("Authorized on account %s", bot.Self.UserName)

    msg := tgbotapi.NewMessage(c.SendMessage, "test ping")

    bot.Send(msg)
}