package slackapi

type Message struct {
    Text     string `json:"text"`
    Channel  string `json:"channel"`
    Username string `json:"username"`
}