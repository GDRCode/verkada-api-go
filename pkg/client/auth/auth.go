// This package is used internally for auth-related requests.
// Functions are exported in case auth tokens or streaming tokens are needed separately.
package auth

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Contains the most recent short-lived auth token and its expiration time.
// All client package methods check this time and request a new token if needed.
type TokenContainer struct {
	Token   string `json:"token"`
	Expires time.Time
}

// Get a short-lived auth token using the Client's key.
func GetAuthToken(key string, baseURL string) (TokenContainer, error) {
	req, _ := http.NewRequest("POST", baseURL+"/token", nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", key)

	ret := TokenContainer{Expires: time.Now().Add(time.Minute * 29)}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return TokenContainer{}, fmt.Errorf("%s - could not retrieve auth token from .env API key, status: %s", err.Error(), resp.Status)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		fmt.Println(err)
		return TokenContainer{}, fmt.Errorf("error parsing GetAuthToken response fields - %s", err.Error())
	}
	return ret, nil
}

// Similar to the normal auth token, but specifically for the Streaming API.
// Internally used to return the response body as a *bytes.Buffer to be parsed into its response struct.
//
// Does attempt to parse just the jwt string and return it if needed externally
func GetStreamingToken(key string, baseURL string) (*bytes.Buffer, string, error) {
	req, _ := http.NewRequest("GET", baseURL+"/cameras/v1/footage/token", nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("%v - could not retrieve streaming token from API key, status: %s", err, resp.Status)
	}

	defer resp.Body.Close()
	var buf bytes.Buffer
	rec := struct {
		Jwt string `json:"jwt"`
	}{}
	tee := io.TeeReader(resp.Body, &buf)
	decode := json.NewDecoder(tee)
	err = decode.Decode(&rec)
	if err != nil {
		return nil, "", fmt.Errorf("%v - could not marshal streaming token into receiver struct", err.Error())
	}
	return &buf, rec.Jwt, nil
}

// Parse the ".env" file in the current working directory.
// File should be formatted as one key:value pair per line, e.g. API_KEY:key_value_here
//
// Used internally to retrieve API key if supplied (must have key "API_KEY" to be recognized)
func GetEnvFromFile() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ":")
		os.Setenv(items[0], items[1])
	}
	return nil
}
