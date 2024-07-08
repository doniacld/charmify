package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/doniacld/charmify/pkg/client"
	"github.com/doniacld/charmify/pkg/habittracker"
)

func main() {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	habitsClient, err := client.New("localhost:28710")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	model := habittracker.New(habitsClient)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
