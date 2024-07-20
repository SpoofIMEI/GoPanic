package panic

import (
	"bytes"
	"fmt"
	"gopanic/internal/config"
	"gopanic/internal/messages"
	"gopanic/internal/presets"
	"os"
	"os/exec"
	"strconv"
)

var TripReady bytes.Buffer

func Trip() {
	for key, val := range config.Instructions.Presets {
		var err error
		switch key {
		case "sdelete":
			err = presets.SecureDelete(val.(string))
		case "shutdown":
			err = presets.Shutdown()
		case "kill":
			err = presets.Kill(val.(string))
		}
		if err != nil {
			messages.Error(err.Error())
		}
	}

	cmd := exec.Command("powershell")
	cmd.Stdin = &TripReady
	cmd.Run()
	os.Exit(0)
}

func Deploy() {
	tripMe := make(chan bool, 1)
	if config.Config["ask_password"] == "false" {
		Trip()
	}
	timeout, _ := strconv.Atoi(config.Config["password_timeout"])
	if timeout > 0 {
		messages.PasswordPrompt(timeout, tripMe, fmt.Sprintf("You have %d seconds to type in the correct password.", timeout))
	} else {
		messages.PasswordPrompt(timeout, tripMe, "Please enter password.")
	}
	if <-tripMe {
		Trip()
	}
}
