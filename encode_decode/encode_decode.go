package encode_decode

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

// GOB

func EncodeGob(v any) []byte {
	var buf bytes.Buffer
	encdr := gob.NewEncoder(&buf)

	if err := encdr.Encode(v); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func DecodeGob(b []byte, res any) {
	buf := bytes.NewBuffer(b)
	dcdr := gob.NewDecoder(buf)

	if err := dcdr.Decode(res); err != nil {
		panic(err)
	}
}

// JSON

func EncodeJSON(v any) []byte {
	res, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return res
}

func DecodeJSON(b []byte, res any) {
	if err := json.Unmarshal(b, res); err != nil {
		panic(err)
	}
}

// XML

func EncodeXML(v any) []byte {
	res, err := xml.Marshal(v)
	if err != nil {
		panic(err)
	}

	return res
}

func DecodeXML(b []byte, res any) {
	if err := xml.Unmarshal(b, res); err != nil {
		panic(err)
	}
}

// YAML

func EncodeYAML(v any) []byte {
	res, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}

	return res
}

func DecodeYAML(b []byte, res any) {
	if err := yaml.Unmarshal(b, res); err != nil {
		panic(err)
	}
}
