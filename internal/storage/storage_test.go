package storage_test

import (
	"os"
	"path"
	"testing"

	tpb "github.com/aaron-g-sanchez/SPOTIFY-CLI/internal/protos"
	"github.com/aaron-g-sanchez/SPOTIFY-CLI/internal/storage"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestLoadToken_Errors(t *testing.T) {
	tempTestDirectory := t.TempDir()

	testStorage := storage.Storage{BaseDirectory: tempTestDirectory}

	testCases := []struct {
		name    string
		storage *storage.Storage
		err     error
	}{{

		name:    "Should not throw error when no file is detected.",
		storage: &testStorage,
		err:     os.ErrNotExist,
	},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := tc.storage.LoadToken()

			if err != nil {
				t.Errorf("Received unexpected error: %v", err)
			}
		})
	}

}

func TestLoadToken_Success(t *testing.T) {
	// TODO: Move this setup to a setup function.
	expectedTestToken := &tpb.TokenData{AccessToken: "test-token", RefreshToken: "test-refresh-token"}

	tempTestDirectory := t.TempDir()

	spotifyDirectoryPath := path.Join(tempTestDirectory, ".spotify-cli")

	if err := os.Mkdir(spotifyDirectoryPath, 0744); err != nil {
		t.Fatalf("Error creating .spotify-cli directory: %v", err)
	}

	tempTokenPath := path.Join(tempTestDirectory, ".spotify-cli", "tokens.json")

	testTokenJSONData, err := protojson.Marshal(expectedTestToken)
	if err != nil {
		t.Fatalf("Error creating JSON: %v", err)
	}

	if err = os.WriteFile(tempTokenPath, testTokenJSONData, 0644); err != nil {
		t.Fatalf("Error writing test token file: %v", err)
	}

	testStorage := storage.Storage{BaseDirectory: tempTestDirectory}

	testCases := []struct {
		name    string
		storage *storage.Storage
		token   *tpb.TokenData
	}{{

		name:    "Should retrieve token.",
		storage: &testStorage,
		token:   expectedTestToken,
	},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			gotToken, err := tc.storage.LoadToken()
			if err != nil {
				t.Fatalf("Unexpected error while loading token: %v", err)
			}

			if !proto.Equal(gotToken, tc.token) {
				t.Errorf("got token: %v, expected token: %v", gotToken, tc.token)

			}

		})
	}

}
