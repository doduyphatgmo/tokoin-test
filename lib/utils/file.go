package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJsonFile(path string, model interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteValue, model)
	if err != nil {
		return err
	}

	defer func() {
		err = jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return nil
}