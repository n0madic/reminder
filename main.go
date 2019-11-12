package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	markName  = "_GO_REMINDER"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hh:mm> <text message>\n", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()

	t, err := time.Parse("2006-01-02 MST 15:04", now.Format("2006-01-02 MST ")+os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if now.After(t) {
		fmt.Println("Set a future time!")
		os.Exit(3)
	}

	diff := t.Sub(now)

	if os.Getenv(markName) == markValue {
		time.Sleep(diff)

		err = beeep.Alert("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}

		fmt.Println("Reminder will be displayed after", diff.Truncate(time.Second))
		os.Exit(0)
	}
}
