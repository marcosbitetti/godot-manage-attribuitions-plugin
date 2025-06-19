package main

import (
	_ "embed"
	"errors"
	"os"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/interfaces/command"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/usecases"
)


func main() {
	argCount := len(os.Args)
	if argCount == 1 {
		println(string(usecases.FormatJSON(nil, errors.New("no command provided"))))
		return
	}

	path, err := infra.ParseDatabasePath(os.Args)
	if err != nil {
		println(string(usecases.FormatJSON(nil, err)))
		return
	}

	storage, err := infra.NewStorage(path)
	if err != nil {
		println(string(usecases.FormatJSON(nil, err)))
		return
	}
	defer storage.CloseDatabase()

	println(string(command.ParseCommand(storage, os.Args)))
}

