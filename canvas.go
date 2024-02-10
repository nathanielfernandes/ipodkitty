package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var canvasEndpoint = "https://canvas-api.ncp.nathanferns.xyz/gen_url/" + os.Getenv("CANVAS_TOKEN")

var client = &http.Client{}

type RunPayload struct {
	Size   []int         `json:"size"`
	Files  []File        `json:"files"`
	Assets []interface{} `json:"assets"`
}

type File struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ImageAsset struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LiteralAsset struct {
	Name    string `json:"name"`
	Literal string `json:"literal"`
}

func GetCanvasUrl(payload *RunPayload) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	body := bytes.NewBuffer(jsonPayload)
	req, err := http.NewRequest("POST", canvasEndpoint, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("canvas api returned %d", resp.StatusCode)
	}

	text := new(bytes.Buffer)
	text.ReadFrom(resp.Body)

	return text.String(), nil
}
