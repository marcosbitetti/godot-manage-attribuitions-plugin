package usecases

import (
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
)

func GetTypes(storage *infra.Storage,  _ []string) []byte {
	return FormatJSON(storage.ListTypes())

}
