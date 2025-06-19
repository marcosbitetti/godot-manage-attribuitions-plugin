package usecases

import (
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/domain"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
)

func GetAttribuitions(storage *infra.Storage, args []string) []byte {
	if len(args) < 3 {
		return FormatJSON(nil, NewErrMissingArgument())
	}
	if len(args) < 4 {
		args = append(args, `{"Text":"","Order":"ASC"}`)
	}

	query, err := domain.NewQuery(args[3])
	if err != nil {
		return FormatJSON(nil, err)
	}

	attribuitions, err := storage.FindAttribuitions(query.Order, query.Text)
	if err != nil {
		return FormatJSON(nil, err)
	}
	return FormatJSON(attribuitions, nil)
}
