package main

import (
	"github.com/aaron-g-sanchez/PROTOTYPE/SPOTIFY-CLI/config"
)

// TODO: Add Auth with PKCE to get Spotify access token.
// 1: Create a code verifier
// 2: Create a code challenge from the verifier (SHA256)
// 3: Request user authorization with code challenge
// 4: Upon successful authorization, get the code from the callback url
// 5: Take aforementioned code and request an access token using the code verifier from the step 1
// 6: Upon successful token creation, use the token to request user data from Spotify

// TODO: Create a player and request current state of the users playback.

func main() {
	// TODO: Check for a valid token.
	config.GetToken()
}
