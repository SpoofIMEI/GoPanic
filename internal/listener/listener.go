package listener

import (
	"bytes"
	"fmt"
	keybind "gopanic/config"
	"gopanic/internal/config"
	"gopanic/internal/messages"
	"gopanic/internal/panic"
	"os"
	"os/signal"
	"syscall"

	"github.com/gen2brain/beeep"
	"golang.design/x/hotkey"
)

func Listen() {
	panic.TripReady = bytes.Buffer{}
	for _, command := range config.Instructions.Commands {
		panic.TripReady.Write([]byte(command + "\n"))
	}

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, os.Interrupt)
	go func() {
		<-term
		beeep.Notify("GoPanic stopped", "GoPanic has stopped running", "assets/information.png")
		os.Exit(0)
	}()
	err := keybind.HotKey.Register()
	if err != nil {
		messages.Error(err.Error())
		return
	}
	err = beeep.Notify("GoPanic running", "GoPanic is now running in the background!", "assets/information.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	wait(keybind.HotKey.Keydown())
}

func wait(ch <-chan hotkey.Event) {
	for {
		<-ch
		panic.Deploy()
	}
}
