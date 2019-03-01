package dm

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type DataParser interface {
	Parse(r io.Reader, v interface{}) error
}

type JsonDataModel struct{}

func (d *JsonDataModel) Parse(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(v)
	if err != nil {
		return errors.Wrap(err, "Could not parse json")
	}
	return nil
}
