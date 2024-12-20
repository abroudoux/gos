package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		err := flagMode()
		if err != nil {
			printErrorAndExit(err)
		}

		os.Exit(0)
	}
}

func flagMode() error {
	flag := os.Args[1]

	if flag == "--version" || flag == "-v" {
		latestVersion, err := getLatestRealeaseVersion("abroudoux", "gos")
		if err != nil {
			printErrorAndExit(err)
		}

		println(latestVersion)
	} else if flag == "--help" || flag == "-h" {
		printHelpManual()
	}

	return nil
}

func printErrorAndExit(err error) {
	println(err)
	os.Exit(0)
}

func printHelpManual() {
	fmt.Println("gos is a template to create Go tools")
	fmt.Println("Usage: gos [options]")
	fmt.Printf("  %-20s %s\n", "gos [--help | -h]", "Show this help message")
}