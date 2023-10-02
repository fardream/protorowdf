package protorowdf

import (
	"encoding/json"

	"github.com/goccy/go-yaml"
	"google.golang.org/protobuf/encoding/protojson"
)

func ParseYAML(s []byte) (*ProtoFile, error) {
	r := &ProtoFile{}
	err := yaml.Unmarshal(s, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func ParseJSON(s []byte) (*ProtoFile, error) {
	r := &ProtoFile{}

	err := protojson.Unmarshal(s, r)
	if err == nil {
		return r, nil
	}

	err = json.Unmarshal(s, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
