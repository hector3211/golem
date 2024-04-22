package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	cli "go-cli-csv/src/internal/cli"
	csv "go-cli-csv/src/internal/csv"
	"os"
)

func init() {
	records := csv.GetFile("inventory.csv")
	if _, err := tea.NewProgram(cli.New(records)).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Shutting down...")
	// p := tea.NewProgram(cli.InitialModel(), tea.WithAltScreen())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Alas, there's been an error: %v", err)
	// 	os.Exit(1)
	// }
}
