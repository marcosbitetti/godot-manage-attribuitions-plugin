package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	const baseUrl = "http://localhost:10010/"

	tempDir := t.TempDir()
	databasePath := tempDir + "/nonexistent.db"
	os.Setenv("DATABASE_PATH", databasePath)

	go main()

	waitServer(t, baseUrl)

	t.Run("should return help text", func(t *testing.T) {
		_, responseBody := makeRequest(t, baseUrl+"help", http.StatusOK)

		const expected = `Usage:`
		if !strings.Contains(responseBody, expected) {
			t.Errorf("Expected response body to be '%s', got %s", expected, responseBody)
		}
	})

	t.Run("should return listTypes", func(t *testing.T) {
		_, responseBody := makeRequest(t, baseUrl+"listTypes", http.StatusOK)

		const expected = `3D Model`
		if !strings.Contains(responseBody, expected) {
			t.Errorf("Expected response body to be '%s', got %s", expected, responseBody)
		}
	})

}

func waitServer(t *testing.T, url string) {
	var err error
	for i := 0; i < 5; i++ {
		_, err = http.Get(url)
		if err == nil {
			break
		}
		if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "no such host") {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		t.Fatalf("Failed to make GET request: %v", err)
	}
	if err != nil {
		t.Fatalf("Failed to make GET request after retries: %v", err)
	}
}

func makeRequest(t *testing.T, url string, expectedResponseCode int) (*http.Response, string) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	bodyString := string(bodyBytes)

	if resp.StatusCode != expectedResponseCode {
		t.Errorf("Expected status %d, got %d, responaw body: %s", expectedResponseCode, resp.StatusCode, bodyString)
	}
	return resp, bodyString
}
