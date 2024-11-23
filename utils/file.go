package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func ReadJSONFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	return json.Unmarshal(byteValue, v)
}

func WriteJSONFile(filePath string, v interface{}) error {
	byteValue, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, byteValue, 0644)
}

func GetCurrentTime() string {
	return fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
}
