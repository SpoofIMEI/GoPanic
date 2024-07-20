package main

import (
	"gopanic/internal/config"
	"gopanic/internal/listener"
	"gopanic/internal/messages"
	"time"
)

func main() {
	setup()
	time.Sleep(5 * time.Second)
}

var defaults = map[string]string{
	"password_timeout": "20",
	"ask_password":     "true",
}

func setup() {
	err := config.Init(nil, defaults)
	if err != nil {
		messages.Error(err.Error())
		return
	}
	listener.Listen()
}
