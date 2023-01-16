package repository

import (
	"encoding/json"
	"errors"
	"os"
)

func appendlocalpost(post Post) (bool, error) {
	file, err := os.OpenFile("./data/post", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {

		return false, errors.New("Could not open post")
	}

	defer file.Close()
	b, err := json.Marshal(post)
	_, err2 := file.WriteString(string(b) + "\n")

	if err2 != nil {

		return false, errors.New("Could not write text to post")

	}
	return true, nil
}
