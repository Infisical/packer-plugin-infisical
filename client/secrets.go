package infisicalclient

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (client Client) GetSecretsRawV3(request GetRawSecretsV3Request) (GetRawSecretsV3Response, error) {
	var secretsResponse GetRawSecretsV3Response

	httpRequest := client.Config.HttpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", USER_AGENT).
		SetQueryParams(map[string]string{
			"environment":            request.Environment,
			"workspaceId":            request.WorkspaceId,
			"expandSecretReferences": strconv.FormatBool(request.ExpandSecretReferences),
		})

	if request.SecretPath != "" {
		httpRequest.SetQueryParam("secretPath", request.SecretPath)
	}

	response, err := httpRequest.Get("api/v3/secrets/raw")

	if err != nil {
		return GetRawSecretsV3Response{}, fmt.Errorf("CallGetSecretsRawV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		// Attempt to extract the specific error message from the JSON response body
		var apiErrorDetail struct {
			Message string `json:"message"`
		}

		if unmarshalErr := json.Unmarshal(response.Body(), &apiErrorDetail); unmarshalErr == nil && apiErrorDetail.Message != "" {
			return GetRawSecretsV3Response{}, fmt.Errorf("CallGetSecretsRawV3: API error: %s", apiErrorDetail.Message)
		}

		// Fallback if JSON parsing fails or the "message" field is not found or is empty.
		return GetRawSecretsV3Response{}, fmt.Errorf("CallGetSecretsRawV3: Unsuccessful response (Status: %s). Please make sure your secret path, workspace and environment name are all correct. Raw body: %s", response.Status(), response.String())
	}

	return secretsResponse, nil
}

func (client Client) GetRawSecrets(secretFolderPath string, envSlug string, workspaceId string) ([]RawV3Secret, error) {
	request := GetRawSecretsV3Request{
		Environment:            envSlug,
		WorkspaceId:            workspaceId,
		ExpandSecretReferences: true,
	}

	if secretFolderPath != "" {
		request.SecretPath = secretFolderPath
	}

	secrets, err := client.GetSecretsRawV3(request)

	if err != nil {
		return nil, err
	}

	return secrets.Secrets, nil
}
