package main

import (
    "flag"
    "github.com/d-kunin/slackapi"
    "github.com/d-kunin/playfeedback"
    "fmt"
)

var arg_webhook string
var arg_text    string
var arg_chat    string
var arg_user    string
var arg_file    string

func init() {
    flag.StringVar(&arg_webhook, "hook", "", "Slack Webhook URL")
    flag.StringVar(&arg_text, "text", "hello", "Message text")
    flag.StringVar(&arg_chat, "chat", "trash", "Chat to send message to")
    flag.StringVar(&arg_user, "user", "Benber", "Message author")
    flag.StringVar(&arg_file, "file", "", "CSV data source")
}

func main() {
    flag.Parse()
    
    if len(arg_webhook) == 0 {
        panic("webhook must be set")
    }
    
    slack := slackapi.NewSlackApi(arg_webhook)
    
    if len(arg_file) > 0 {
        fmt.Println("Sending reviews from file")
        rx, err := playfeedback.FromCsvFile(arg_file)
        if err != nil {
            fmt.Printf("BAD THING HAPPENED: %v\n", err)
            return
        }
        rx = playfeedback.FilterRecent(rx)
        fmt.Printf("Number of recent reviews: %d\n", len(rx))
    
        count := 0
        for _,r := range rx {
            if len(r.Text) > 10 {
                slackmsg := playfeedback.ReviewToSlackMessage(&r)
                slackmsg.Channel = arg_chat
                slack.SendMessage(slackmsg)
                count += 1
            }
        }
        fmt.Printf("Send %d new reviews\n", count)
    } else {
        fmt.Println("Sending simple message")
        slack.SendMessage(&slackapi.Message{
            arg_text,
            arg_chat,
            arg_user,
        })
    }
}
