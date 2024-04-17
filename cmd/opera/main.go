package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"os"

	"github.com/Fantom-foundation/go-opera/cmd/opera/launcher"
)

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
