package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var defaultHTTPClient = &http.Client{
	Timeout: 30 * time.Second,
}

func HTTPGet(url string) (*http.Response, error) {
	return defaultHTTPClient.Get(url)
}

func HTTPGetOK(url string) ([]byte, error) {
	resp, err := HTTPGet(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request to %s failed (status %d)", url, resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func DecodeJSONFromURL(url string, dest interface{}) error {
	data, err := HTTPGetOK(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

func DownloadFile(filePath, url string) error {
	resp, err := HTTPGet(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download from %s failed (status %d)", url, resp.StatusCode)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}
