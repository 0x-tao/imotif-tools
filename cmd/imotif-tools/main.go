package main

import (
	"fmt"
	"log"
	"os"

	"github.com/imotif-tools/internal/ai"
	"github.com/imotif-tools/internal/cli"
	"github.com/imotif-tools/internal/config"
	"github.com/imotif-tools/internal/odoo"
	"github.com/imotif-tools/internal/update"
)

const version = "v1.1.3"

func main() {
	args := os.Args

	if len(args) == 1 {
		// Init Banner
		banner := cli.NewBanner()
		fmt.Println("imotif-tools version:", version)
		banner.PrintBanner()
		return
	}

	switch args[1] {
	case "--version", "-v":
		fmt.Println("imotif-tools version:", version)
		return
	case "--help", "-h":
		// Init Banner
		banner := cli.NewBanner()
		banner.PrintHelp()
		return
	case "test":
		odoo := odoo.NewTester(args)
		if err := odoo.RunTest(); err != nil {
			log.Fatal("Failed to run test: ", err)
		}
		return
	case "commit":
		prompter := cli.NewCommitPrompter(args)
		if err := prompter.RunCommit(); err != nil {
			log.Fatal("Failed to run commit prompt: ", err)
		}
		return
	case "init":
		// Init AliasInitializer
		initializer := cli.NewAliasInitializer()
		initializer.Run()
		return
	case "update":
		// Init SelfUpdater
		selfUpdater, err := update.NewSelfUpdater()
		if err != nil {
			log.Fatal("Failed to create self updater: ", err)
		}
		fmt.Println("Updating imotif-tools...")
		selfUpdater.Run()
		return
	case "magic":
		cfg := config.Load()
		ai := ai.NewGeminiProvider(cfg.Model, cfg.APIKey)
		if err := ai.RunCommand(); err != nil {
			log.Fatal("Failed to run AI command: ", err)
		}
		return
	default:
		fmt.Println("Invalid command.")
		return
	}
}
