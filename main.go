package main

import (
	"flag"
	"fmt"
)

func main() {
	basic := flag.Bool("basic", false, "Use Markdown Basic(Markdown Common by default).")
	help := flag.Bool("help", false, "Print usage message.")
	port := flag.Int("port", 6060, "Port to listen.")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	args := flag.Args()
	if len(args) != 1 {
		printError("No file provided. Please check the usage.\n")
		printHelp()
		return
	}

	orange := NewOrange(args[0])

	if *basic {
		orange.UseBasic()
	}

	orange.Run(*port)
}

func printError(message string) {
	fmt.Printf("\x1b[31m%s\x1b[39;49m\n", message)
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  orange [options] file\n")
	fmt.Println("Options:")
	flag.PrintDefaults()
}
