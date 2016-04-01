package main

import (
	"fmt"
	"strconv"
)

var completionsTopic = &Topic{
	Name:        "completions",
	Description: "generate shell tab completions",
	Hidden:      true,
}

var completionsCmd = &Command{
	Topic:        "completions",
	Hidden:       true,
	VariableArgs: true,
	Flags:        []Flag{{Name: "cword", Required: true, HasValue: true}},
	Run: func(ctx *Context) {
		SetupBuiltinPlugins()
		cli.LoadPlugins(GetPlugins())
		cword, _ := strconv.Atoi(ctx.Flags["cword"].(string))
		args := ctx.Args.([]string)
		opts := []string{}
		switch cword {
		case 1:
			// commands
			for _, command := range cli.Commands {
				if command.Hidden {
					continue
				}
				if command.Command == "" {
					opts = append(opts, fmt.Sprintf("%s\n", command.Topic))
				} else {
					opts = append(opts, fmt.Sprintf("%s:%s\n", command.Topic, command.Command))
				}
			}
		default:
			//cur := args[cword]
			prev := args[cword-1]
			switch prev {
			case "--app":
				apps, err := apps()
				ExitIfError(err, false)
				for _, app := range apps {
					opts = append(opts, app.Name)
				}
			}
		}

		for _, opt := range opts {
			Println(opt)
		}
	},
}
