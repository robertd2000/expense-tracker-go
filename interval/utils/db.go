package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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

func Delete(filePath string) error {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return errors.New("cannot delete file: file does not exist")
    }

    err := os.Remove(filePath)
    if err != nil {
        return fmt.Errorf("unable to delete file: %w", err)
    }

    return nil
}
