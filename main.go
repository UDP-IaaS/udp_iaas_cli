package main

import (
	"fmt"
	"log"

	"udp_iaas_cli/shell"
	"udp_iaas_cli/utils"

	"github.com/c-bata/go-prompt"
)


func main() {
	// Connection test
	if err := utils.TestGrpcConnection(); err != nil {
		log.Fatalf("Connection to local daemon failed: %v", err)
	}

	fmt.Println("Welcome to UDP IaaS Interactive Shell")
	fmt.Println("Type 'help' to see available commands")

	// Interactive shell, check /shell
	p := prompt.New(
		shell.Executor,
		shell.Completer,
		prompt.OptionTitle("Interactive Shell"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)

	p.Run()
}