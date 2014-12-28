package slackapi

import (
    "net/http"
    "fmt"
    "net/url"
    "encoding/json"
)

type Slack interface {
    SendMessage(message * Message)
}

type SlackApi struct {
    webhook string
}

func NewSlackApi(webhook string) * SlackApi {
    return &SlackApi{webhook}
}

func (slack * SlackApi) SendMessage(message * Message) {
    json,_ := json.Marshal(message)
    jsonstr := string(json)
    vals := url.Values{"payload": {jsonstr}}
    
    client := &http.Client{}
    resp, err := client.PostForm(
        slack.webhook,
        vals,
    )
    
    if err != nil {
        fmt.Printf("Error sending message: %v\n", err)
    } else {
        fmt.Printf("Message sent: %v\n", resp)
    }
}