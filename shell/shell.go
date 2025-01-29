package shell

import (
	"fmt"
	"strings"

	"udp_iaas_cli/utils"

	"github.com/c-bata/go-prompt"
)


var Commands = []string{
    "version",
	"help",
	"exit",
}


func Completer(d prompt.Document) []prompt.Suggest {
    var suggestions []prompt.Suggest
    
    for _, cmd := range Commands {
        suggestions = append(suggestions, prompt.Suggest{
            Text:        cmd,
            Description: fmt.Sprintf("Execute %s command", cmd),
        })
    }

    // 현재 입력된 텍스트
    word := d.GetWordBeforeCursor()
    
    // 입력된 텍스트에 맞는 suggestion들만 필터링
    return prompt.FilterHasPrefix(suggestions, word, true)
}

// Executor processes and executes user input commands.
// If an unknown command is entered, it displays an error message.
func Executor(input string) {
    input = strings.TrimSpace(input)
    
    switch input {
    case "help":
        fmt.Println("Available Commands:")
        for _, cmd := range Commands {
            fmt.Printf("  %s\n", cmd)
        }

    case "exit":
        Exit()

    case "version":
        utils.Version()
        fmt.Println("Shell Client version: 0.0.1")
    
    default:
        if input != "" {
            fmt.Printf("Unknown command: %s\n", input)
        }
    }
}