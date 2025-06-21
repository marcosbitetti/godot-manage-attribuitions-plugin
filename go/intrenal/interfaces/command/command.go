package command

import (
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/usecases"
)

const helpCommand = "help"

func ParseCommand(storage *infra.Storage, args []string) string {
	commands := usecases.Commands()
	if len(args) < 3 {
		return string(commands[helpCommand](nil, nil))
	}

	caller, has := commands[args[2]]
	if !has {
		caller = commands[helpCommand]
	}

	return string(caller(storage, args))
}

