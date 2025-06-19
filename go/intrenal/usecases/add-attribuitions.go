package usecases

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/domain"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
)

func AddAttribuition(storage *infra.Storage, args []string) []byte {
	if len(args) < 4 {
		return FormatJSON(nil, NewErrMissingArgument())
	}
	var t domain.Attribuition
	if err := json.Unmarshal([]byte(args[3]), &t); err != nil {
		return FormatJSON(nil, errors.Wrap(err, "invalid type"))
	}
	if t.Name == "" ||
		t.Link == "" ||
		t.Author == "" ||
		t.Type == "" ||
		t.Licence == "" {
		return FormatJSON(nil, NewErrInvalidValue())
	}
	if err := storage.AddAttribuition(t.Name, t.FileName, t.Link, t.Author, t.Type, t.Licence); err != nil {
		return FormatJSON(nil, errors.Wrap(err, "error adding attribuition"))
	}
	return FormatJSON(SuccessMsg, nil)

}
