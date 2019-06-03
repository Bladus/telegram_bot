package main

import (
    "gopkg.in/telegram-bot-api.v4"
    "application/config"
    "encoding/json"
    "io/ioutil"
    "regexp"
    "flag"
    "fmt"
    "log"
    "os"
)

const (
    TARGET = "https://ria.ru/photolents"
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

    //msg := tgbotapi.NewMessage(c.SendMessage, "test ping")

    //bot.Send(msg)

    response, err := request(TARGET)
    defer response.Body.Close()

    if err != nil {
        log.Printf("ERROR request(TARGET) :: %v\n", err.Error())
    }

    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Printf("ERROR ioutil.ReadAll(response.Body) :: %v\n", err.Error())
    }

    re := regexp.MustCompile(`<a[^>]+\bhref=["']([^"']+\.html)" class="list-item__image">`)
    items := re.FindAllStringSubmatch(string(content), -1)

    stamp := items[0][1]

    last, err := ioutil.ReadFile("./last.txt")
    check_error(err)
    log.Printf("last url :: %s\n", string(last))

    if(string(last) == stamp) {
        return
    }

    s := []byte(stamp)
    err = ioutil.WriteFile("./last.txt", s, 0644)
    check_error(err)

    response2, err := request(stamp)
    defer response2.Body.Close()
    check_error(err)

    content2, err := ioutil.ReadAll(response2.Body)
    if err != nil {
        log.Printf("ERROR ioutil.ReadAll(response2.Body) :: %v\n", err.Error())
    }

    re2 := regexp.MustCompile(`<img media-type="photo" alt="[^>]+" title="([^>]+)" src="([^>]+)">`)
    images := re2.FindAllStringSubmatch(string(content2), -1)

    cap := regexp.MustCompile(`<div class="article__photo-item-text"><p>([^>]+)</p></div>`)
    captions := cap.FindAllStringSubmatch(string(content2), -1)

    for i, v := range images {
        fmt.Printf("%s\n\n", captions[i][1])
        photo := CreatePhoto(c.SendMessage, captions[i][1], v[1])
        bot.Send(photo)
    }
}