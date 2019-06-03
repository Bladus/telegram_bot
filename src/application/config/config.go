package config

type Config struct {
    Telegram struct {
        Bot   string `json:"bot"`
        Token string `json:"token"`
    }
    SendMessage int64 `json:"send_message"`
}