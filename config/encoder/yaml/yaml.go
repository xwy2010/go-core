package yaml

import (
	"github.com/ghodss/yaml"
	"github.com/xwy2010/go-core/config/encoder"
)

type yamlEncoder struct{}

func (y yamlEncoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (y yamlEncoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

func (y yamlEncoder) String() string {
	return "yaml"
}

func NewEncoder() encoder.Encoder {
	return yamlEncoder{}
}
