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

type Response struct {
	Error   string `json:"error"`
	Path    string `json:"path"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func executeRequest(payload any, operationName string, logErrors bool) error {
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

	var response Response
	err = json.Unmarshal(resb, &response)
	if err != nil {
		return fmt.Errorf("unable to unmarshal response body: %w", err)
	}

	if response.Message == "success" {
		log.Info("Successfully executed operation", "operation", operationName)
	}

	if logErrors && response.Code != "" {
		log.Errorf(
			"Operation %s failed: Received %s error for path (%s): %s",
			operationName, response.Code, response.Path, response.Error,
		)
	}

	return nil
}
