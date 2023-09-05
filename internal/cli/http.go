package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

type TableName struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

func executeRequest(payload any) error {
	client := &http.Client{}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("unable to marshal http request: %w", err)
	}

	if dryRun {
		log.Debug("dry run... sending request", "body", string(body))
		return nil
	}

	log.Debug("sending request", "body", string(body))

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("unable to build http request: %w", err)
	}

	req.Header["x-hasura-admin-secret"] = []string{adminKey}
	req.Header["X-Hasura-Role"] = []string{"admin"}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("unable to execute http request: %w", err)
	}

	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body: %w", err)
	}

	log.Debug("received response", "body", string(resb))

	return nil
}
