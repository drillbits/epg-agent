package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/drillbits/epg-agent"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK = iota
	ExitCodeError
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		in     string
		out    string
		config string

		showVersion bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&in, "in", "", "Source name or path to a input file.")
	// flags.StringVar(&in, "i", "", "Source name or path to a input file.(Short)")

	flags.StringVar(&out, "out", "", "Destination name.")
	// flags.StringVar(&out, "o", "", "Destination name.(Short)")

	flags.StringVar(&config, "config", "", "Path to a config file.")
	// flags.StringVar(&config, "c", "", "Path to a config file.(Short)")

	flags.BoolVar(&showVersion, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if showVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", name, version)
		return ExitCodeOK
	}

	conf, err := agent.ReadConfigPath(config)
	if err != nil {
		fmt.Fprintf(cli.errStream, "Error: %s", err)
		return ExitCodeError
	}

	// TODO override config
	_ = in
	_ = out

	// TODO run
	_ = conf

	return ExitCodeOK
}
