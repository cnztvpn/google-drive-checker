package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"google.golang.org/api/drive/v3"

	"github.com/whywaita/google-drive-checker/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetClient(c config.Config) *http.Client {
	oconfig, err := google.ConfigFromJSON([]byte(c.CredJson), drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return getClientByConfig(oconfig, c)
}

// Retrieve a token, saves the token, then returns the generated client.
func getClientByConfig(config *oauth2.Config, c config.Config) *http.Client {
	var tok oauth2.Token
	tok = getTokenFromRefreshToken(c.GoogleDriveToken, config)

	return config.Client(context.Background(), &tok)
}

func getTokenFromRefreshToken(refreshToken string, config *oauth2.Config) oauth2.Token {
	tokenURL := "https://www.googleapis.com/oauth2/v4/token"
	values := url.Values{}
	values.Add("refresh_token", refreshToken)
	values.Add("client_id", config.ClientID)
	values.Add("client_secret", config.ClientSecret)
	values.Add("grant_type", "refresh_token")

	resp, err := http.PostForm(tokenURL, values)
	if err != nil {
		log.Fatalf("Unable to POST refresh token %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable parse response body %v", err)
	}

	tok := oauth2.Token{}

	err = json.Unmarshal(body, &tok)
	if err != nil {
		log.Fatalf("Unable to parse refresh token %v", err)
	}
	return tok
}
