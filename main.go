package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	numArgs := len(os.Args) - 1
	if numArgs == 0 {
		fmt.Println(`dark tells your system's appearance preferences to set dark mode to:
			on			dark mode on
			off			dark mode off
			toggle	toggle dark mode`)
	}
	if numArgs < 1 || numArgs > 1 {
		fmt.Println("dark takes one argument: on, off, or toggle")
	} else {
		darkStatus := os.Args[1]
		if darkStatus != "on" && darkStatus != "off" && darkStatus != "toggle" {
			fmt.Println("dark takes one argument: on, off, or toggle")
		} else {
			var cmd *exec.Cmd
			if darkStatus == "on" {
				cmd = exec.Command("osascript", "-e", `tell application "System Events" to tell appearance preferences to set dark mode to true`)
			} else if darkStatus == "off" {
				cmd = exec.Command("osascript", "-e", `tell application "System Events" to tell appearance preferences to set dark mode to false`)
			} else {
				cmd = exec.Command("osascript", "-e", `tell application "System Events" to tell appearance preferences to set dark mode to not dark mode`)
			}
			err := cmd.Run()
			if err != nil {
				log.Printf("Command finished with error: %v", err)
			}
		}
	}
}
