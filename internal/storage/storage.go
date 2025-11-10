package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
)

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoadToken() (*TokenData, error) {
	// TODO: Attempt to read the ~/.spotify-cli/token.json file.
	path, err := getTokenPath()
	if err != nil {
		log.Fatalf("Error retrieving token path:%v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	var token TokenData

	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

/* Returns the path to the token file. */
func getTokenPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".spotify-cli", "tokens.json"), nil
}
