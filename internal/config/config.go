package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type panicInstructions struct {
	Commands []string               `json:"commands"`
	Presets  map[string]interface{} `json:"presets"`
}

var Config = make(map[string]string)

var Instructions = &panicInstructions{}

func Init(mandatory []string, defaults map[string]string) error {
	//Global configs
	Config = defaults
	handle, err := os.Open("config/global.conf")
	if err != nil {
		return err
	}
	defer handle.Close()

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 3 || !strings.Contains(line, "=") {
			continue
		}
		sections := strings.Split(line, "=")
		key, value := sections[0], strings.Join(sections[1:], "=")
		Config[key] = value
	}
	var errorMsg string
	for _, key := range mandatory {
		if Config[key] == "" {
			errorMsg += "setting \"" + key + "\" not defined" + "\n"
		}
	}
	if errorMsg != "" {
		return errors.New(errorMsg)
	}

	//Panic instructions
	panicContent, err := os.ReadFile("config/panic.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(panicContent, Instructions)
	if err != nil {
		return err
	}

	return verify()
}

func verify() error {
	var finalError string
	_, err := strconv.Atoi(Config["password_timeout"])
	if err != nil {
		finalError += "password timeout must be a number (seconds)\n"
	}

	if Config["ask_password"] != "true" && Config["ask_password"] != "false" {
		finalError += "ask_password must be type boolean (true/false)\n"
	}

	for key, val := range Instructions.Presets {
		switch key {
		case "SecureDelete":
			_, k := val.(string)
			if !k {
				return errors.New("argument for securedelete must be type string")
			}
		case "kill":
			_, k := val.(string)
			if !k {
				return errors.New("argument for kill must be type string")
			}
		}
	}

	if finalError != "" {
		return errors.New(finalError)
	}
	return nil
}
