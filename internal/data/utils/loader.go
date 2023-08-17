// internal/data/utils/loader.go

package utils

import (
	"encoding/json"
	"os"
)

// LoadJSONData loads data from a JSON file into the provided interface.
func LoadJSONData(filepath string, data interface{}) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}
