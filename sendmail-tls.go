package main

import (
   "crypto/tls"
    "encoding/json"
    "os"
    
    "gopkg.in/gomail.v2"
)

func getConfig(f string, v interface{}) error {
    fr, err := os.Open(f)
    if err != nil {
        return err
    }
    defer fr.Close()
    return json.NewDecoder(fr).Decode(v)
}

func main() {
    if len(os.Args) != 3 {
        return
    }
    var config struct {
        Host    string   `json:"host"`
        Port    int      `json:"port"`
        User    string   `json:"user"`
        Pass    string   `json:"pass"`
        From    string   `json:"from"`
        To      []string `json:"to"`
        Subject string   `json:"subject"`
    }
    err := getConfig(os.Args[1], &config)
    if err != nil {
        return
    }

    mail := gomail.NewDialer(config.Host, config.Port, config.User, config.Pass)
    mail.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    err = mail.DialAndSend(gomail.NewMessage(func(m *gomail.Message) {
        m.SetHeader("From", config.From)
        m.SetHeader("To", config.To...)
        m.SetHeader("Subject", config.Subject)
        m.SetBody("text/plain", os.Args[2])
    }))
    if err != nil {
        panic(err)
    }
}
