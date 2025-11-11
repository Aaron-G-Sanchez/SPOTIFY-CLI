package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Storage struct {
	BaseDirectory string
}

func (s *Storage) LoadToken() (*TokenData, error) {
	tokenPath := createTokenPath(s.BaseDirectory)

	data, err := os.ReadFile(tokenPath)
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
func createTokenPath(baseDirectory string) string {
	return path.Join(baseDirectory, ".spotify-cli", "tokens.json")
}
