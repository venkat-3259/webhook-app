package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HttpRequest(v any, reqMethod, reqURL, authToken string) (StatusCode int, response string, err error) {

	byteData, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return 500, "", err
	}
	// log.Println("Request", string(byteData))

	req, err := http.NewRequest(reqMethod, reqURL, bytes.NewBuffer(byteData))
	if err != nil {
		log.Println(err)
		return 500, "", err
	}

	if authToken != "" {
		req.Header.Add("Authorization", authToken)
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return 500, "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode, string(body), err

}
