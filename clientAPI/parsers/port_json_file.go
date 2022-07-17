package parsers

import (
	"encoding/json"
	"fmt"
	"io"
)

// PortFile represents the Json structure on the file
type PortFile struct {
	Key         string
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func ParseJsonPortFile(stream io.ReadCloser) (func() (*PortFile, error), error) {
	dec := json.NewDecoder(stream)

	err := checkToken(dec, '{')
	if err != nil {
		return nil, err
	}

	// return a closure func to interact with the stream
	return func() (*PortFile, error) {
		return parseObject(dec)
	}, nil
}

func checkToken(dec *json.Decoder, token json.Delim) error {
	t, err := dec.Token()
	if err != nil {
		return err
	}

	delim, ok := t.(json.Delim)
	if !ok || delim != token {
		return fmt.Errorf("unexpected token: %s", string(token))
	}

	return nil
}

// function that will parse objects of port file
func parseObject(dec *json.Decoder) (*PortFile, error) {
	if !dec.More() {
		return nil, nil
	}
	key, err := dec.Token()
	if err != nil {
		return nil, err
	}

	port := new(PortFile)
	err = dec.Decode(&port)
	if err != nil {
		return nil, err
	}
	keyStr, ok := key.(string)
	if !ok {
		return nil, fmt.Errorf("failed to parse key: %v", key)
	}
	port.Key = keyStr
	return port, nil
}
