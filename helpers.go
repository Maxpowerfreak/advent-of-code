package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetInputResponseBody(year, day int) (io.ReadCloser, error) {
	cookie, err := readCookie()
	if err != nil {
		return nil, fmt.Errorf("error reading: %x", err)
	}

	req, err := buildRequest(year, day, cookie)
	if err != nil {
		return nil, fmt.Errorf("error reading request: %x", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing request: %x", err)
	}

	return resp.Body, nil
}

func readCookie() (string, error) {
	data, err := os.ReadFile("cookie")
	if err != nil {
		// vsc debugger quick hack
		data, err = os.ReadFile("../../../cookie")
		if err != nil {
			// hacky once more, but we might not be in Part1/2 so check one level above
			data, err = os.ReadFile("../../cookie")
			if err != nil {
				return "", fmt.Errorf("unable to read cookie file: %x", err)
			}
		}
	}

	return string(data), nil
}

func buildRequest(year, day int, cookie string) (*http.Request, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %x", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})

	return req, nil
}
