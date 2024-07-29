package dm

import (
	"encoding/json"
	"io"
	"sigs.k8s.io/yaml"

	"github.com/pkg/errors"
)

type DataParser interface {
	Parse(r io.Reader, v interface{}) error
}

type YamlDataModel struct{}

func (d *YamlDataModel) Parse(r io.Reader, v interface{}) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = yaml.Unmarshal(b, v)
	if err != nil {
		return errors.Wrap(err, "Could not parse yaml")
	}
	return nil
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
