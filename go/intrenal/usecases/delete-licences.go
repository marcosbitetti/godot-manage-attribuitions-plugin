package usecases

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/domain"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
)

func DeleteLicence(storage *infra.Storage, args []string) []byte {
	if len(args) < 4 {
		return FormatJSON(nil, NewErrMissingArgument())
	}
	var t domain.Type
	if err := json.Unmarshal([]byte(args[3]), &t); err != nil {
		return FormatJSON(nil, errors.Wrap(err, "invalid type"))
	}
	if t.Id == 0 {
		return FormatJSON(nil, NewErrInvalidValue())
	}
	if err := storage.DeleteLicence(t.Id); err != nil {
		return FormatJSON(nil, errors.Wrap(err, "error deleting licence"))
	}
	return FormatJSON(SuccessMsg, nil)

}
