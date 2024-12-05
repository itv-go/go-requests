package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get[T any](url string, obj *T) (*T, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
