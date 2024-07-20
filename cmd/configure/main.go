package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var config string
	fmt.Println("    GoPanic setup")
	fmt.Print(`
modes:
 1 - Immediately execute the panic action
 2 - When panic hotkey is pressed, ask for password to disarm it
 `)
	var mode int
	var err error
	for {
		fmt.Print("select mode:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		mode, err = strconv.Atoi(scanner.Text())
		if err != nil || mode > 2 || mode < 1 {
			fmt.Println("invalid mode! valid modes: 1,2")
			continue
		}
		if mode == 1 {
			config += "ask_password=false\n"
		} else {
			config += "ask_password=true\n"
		}
		break
	}

	if mode == 2 {
		var password string
		for {
			fmt.Print("new password:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			pass1 := scanner.Text()
			fmt.Print("repeat password:")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			pass2 := scanner.Text()
			if pass1 != pass2 {
				fmt.Println("password don't match! let's try that again :)")
				continue
			}
			password = pass1
			break
		}
		passHash, err := bcrypt.GenerateFromPassword([]byte(password), -1)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(0)
		}
		config += "password=" + string(passHash) + "\n"

		fmt.Print("how many seconds should the user have time to enter the password? (0 for infinite amount of time)\n:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		secs, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(0)
		}

		config += "password_timeout=" + strconv.Itoa(secs) + "\n"
	}
	handle, err := os.Create("config/global.conf")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(0)
	}
	handle.WriteString(config)
}
