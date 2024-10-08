package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfiguration[C any](filepath string) (*C, error) {
	conf := new(C)
	def, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(def, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
