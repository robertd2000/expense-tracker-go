package utils

import (
	"encoding/json"
	"os"
)

func SerializeToJSON[T any](data T) ([]byte, error) {
	res, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return nil, err
	}

	return res, nil
}

func SaveToJSON(filename string, task []byte) error {
	err := os.WriteFile(filename ,task ,0755)

	if err != nil {
		return err
	}

	return nil
}

func ReadFromJSON(filename string) ([]byte, error) {
	res, err := os.ReadFile(filename)	

	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeserializeFromJSON[T any](data []byte) (T, error) {
	var result T

	if err := json.Unmarshal(data, &result); err != nil {
		return result, err
	}

	return result, nil
}
