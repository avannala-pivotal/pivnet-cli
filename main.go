package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/pivotal-cf/pivnet-cli/commands"
	"github.com/pivotal-cf/pivnet-cli/errorhandler"
	"github.com/pivotal-cf/pivnet-cli/version"
)

var (
	// buildVersion is deliberately left uninitialized so it can be set at compile-time
	buildVersion string
)

func main() {
	if buildVersion == "" {
		version.Version = "dev"
	} else {
		version.Version = buildVersion
	}

	parser := flags.NewParser(&commands.Pivnet, flags.HelpFlag)

	if len(os.Args) > 1 && os.Args[1] == "man" {
		parser.WriteManPage(os.Stdout)
		return
	}

	_, err := parser.Parse()
	if err != nil {
		if err == commands.ErrShowHelpMessage {
			helpParser := flags.NewParser(&commands.Pivnet, flags.HelpFlag)
			helpParser.NamespaceDelimiter = "-"
			_, _ = helpParser.ParseArgs([]string{"-h"})
			helpParser.WriteHelp(os.Stdout)
			os.Exit(0)
		}

		// Do not consider the built-in help an error
		if e, ok := err.(*flags.Error); ok {
			if e.Type == flags.ErrHelp {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(0)
			}
		}

		if err == errorhandler.ErrAlreadyHandled {
			os.Exit(1)
		}

		coloredMessage := fmt.Sprintf(errorhandler.RedFunc(err.Error()))
		fmt.Fprintln(os.Stderr, coloredMessage)
		os.Exit(1)
	}
}
