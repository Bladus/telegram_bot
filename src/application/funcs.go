package main

import (
    "gopkg.in/telegram-bot-api.v4"
    "net/http"
    "log"
)

func CreatePhoto(chat_id int64, caption, link string) tgbotapi.PhotoConfig {
    photo := tgbotapi.NewPhotoShare(chat_id, link)
    photo.Caption = caption

    return photo
}

func request(path string) (*http.Response, error) {
    client := &http.Client{}
    log.Printf("[GET] %v\n", path)

    req, err := http.NewRequest("GET", path, nil)
    if err != nil {
        log.Printf("ERROR http.NewRequest :: %v\n", err.Error())
    }

    return client.Do(req)
}

func check_error(err error) {
    if err != nil {
        log.Println(err)
    }
}