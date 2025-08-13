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

type TokenContainer struct {
	Token   string `json:"token"`
	Expires time.Time
}

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
