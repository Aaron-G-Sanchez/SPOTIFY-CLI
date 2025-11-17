package storage

import (
	"errors"
	"os"
	"path"

	tpb "github.com/aaron-g-sanchez/SPOTIFY-CLI/internal/protos"
	"google.golang.org/protobuf/encoding/protojson"
)

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Storage struct {
	BaseDirectory string
}

func (s *Storage) LoadToken() (*tpb.TokenData, error) {

	tokenPath := createTokenPath(s.BaseDirectory)

	data, err := os.ReadFile(tokenPath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	token := tpb.TokenData{}

	if err := protojson.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

/* Returns the path to the token file. */
func createTokenPath(baseDirectory string) string {
	return path.Join(baseDirectory, ".spotify-cli", "tokens.json")
}
