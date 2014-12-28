package playfeedback

import (
    "github.com/d-kunin/slackapi"
    "strconv"
    "fmt"
    "strings"
)

func ReviewToSlackMessage(r * Review) *slackapi.Message {
    rate    := strconv.Itoa(r.Rating)
    version := strconv.Itoa(r.Version)
    title := "User Review"
    if len(r.Title) > 0 {
        title = r.Title
    }
    title = rate + " " + title
    
    msgParts := []string {
        Linkify(r.Link, title),
        Linkify("http://www.google.com/search?q=" + r.Device, r.Device),
        "Version: " + version,
        r.Text,
    }
    msg := slackapi.Message {
        strings.Join(msgParts, "\n"),
        "",
        title,
    }
    return &msg
}

func Linkify(url, text string) string {
    return fmt.Sprintf("<%s|%s>", url, text)
}