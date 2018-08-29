package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetJSONFromURL does exactly what the name suggests. Takes in an URL, returns
//   a map containing the JSON.
func GetJSONFromURL(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("GET failed")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Read failed")
	}

	var decoded map[string]interface{}
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return nil, errors.New("JSON Unmarshal failed")
	}

	return decoded, nil
}

// GetKeyFromJSON takes in a map containing a JSON, and searches it for key
func GetKeyFromJSON(json map[string]interface{}, key string, recursive bool) (string, error) {
	for k, v := range json {
		if mv, ok := v.(map[string]interface{}); ok {
			if recursive {
				GetKeyFromJSON(mv, key, recursive)
			}
		} else {
			if k == key {
				return fmt.Sprintf("%v", v), nil
			}
		}
	}
	return "", errors.New("Key not found")
}
