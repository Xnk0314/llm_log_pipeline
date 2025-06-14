package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeRequest[T any](method, url, authorizationKey string, body io.Reader) (*T, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorizationKey))

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var t T
	err = json.Unmarshal(responseBytes, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
