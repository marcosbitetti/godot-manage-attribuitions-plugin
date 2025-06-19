package usecases

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/domain"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
)

func UpdateAttribuition(storage *infra.Storage, args []string) []byte {
	if len(args) < 4 {
		return FormatJSON(nil, NewErrMissingArgument())
	}
	var t domain.Attribuition
	if err := json.Unmarshal([]byte(args[3]), &t); err != nil {
		return FormatJSON(nil, errors.Wrap(err, "invalid type"))
	}
	if t.Name == "" || t.Link == "" || t.Author == "" || t.FileName == "" ||
		t.Type == "" || t.Licence == "" ||
		t.Id == 0 {
		return FormatJSON(nil, NewErrInvalidValue())
	}
	if err := storage.UpdateAttribuition(t.Id, t.Name, t.FileName, t.Author, t.Link, t.Type, t.Licence); err != nil {
		return FormatJSON(nil, errors.Wrap(err, "error updating attribuition"))
	}
	return FormatJSON(SuccessMsg, nil)

}
