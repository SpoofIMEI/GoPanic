package messages

import (
	"gopanic/internal/config"
	"time"

	"github.com/martinlindhe/inputbox"
	"github.com/tawesoft/golib/v2/dialog"
	"golang.org/x/crypto/bcrypt"
)

var prefix = "GoPanic:\n"

func Info(msg string) {
	go dialog.Info(prefix + msg)
}
func Error(msg string) {
	go dialog.Error(prefix + msg)
}
func Question(msg string) (bool, error) {
	return dialog.Ask(prefix + msg)
}

func PasswordPrompt(timeout int, tripMe chan bool, msg string) {
	if timeout > 0 {
		go func() {
			time.Sleep(time.Duration(timeout) * time.Second)
			tripMe <- true
		}()
	}

	pass, ok := inputbox.InputBox("Enter password", msg, "")
	if ok {
		err := bcrypt.CompareHashAndPassword([]byte(config.Config["password"]), []byte(pass))
		if err != nil {
			tripMe <- true
		} else {
			Info("disarmed")
			tripMe <- false
		}
	} else {
		tripMe <- true
	}
}
