package domain

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func NewAttribuition(id int64, name string, fileName string, typeName string,
	author string, link string, licence string, licenceUrl string) *Attribuition {
	return &Attribuition{
		Id:         id,
		Name:       name,
		FileName:   fileName,
		Type:       typeName,
		Author:     author,
		Link:       link,
		Licence:    licence,
		LicenceUrl: licenceUrl,
	}
}

func NewType(id int64, name string) *Type {
	return &Type{
		Id:   id,
		Name: name,
	}
}

func NewLicence(id int64, name string, link string) *Licence {
	return &Licence{
		Id:   id,
		Name: name,
		Link: link,
	}
}

func NewQuery(raw string) (*Query, error) {
	q := Query{}
	if err := json.Unmarshal([]byte(raw), &q); err != nil {
		return nil, errors.Wrap(err, "cant unmarshal query")
	}
	if q.Order != "ASC" && q.Order != "DESC" {
		q.Order = "ASC"
	}
	return &q, nil
}
