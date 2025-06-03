package infisicalclient

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (client Client) GetSingleSecretByIDV3(request GetSingleSecretByIDV3Request) (GetSingleSecretByIDV3Response, error) {
	var secretsResponse GetSingleSecretByIDV3Response
	response, err := client.Config.HttpClient.
		R().
		SetResult(&secretsResponse).
		SetHeader("User-Agent", USER_AGENT).
		Get(fmt.Sprintf("api/v3/secrets/raw/id/%s", request.ID))

	if err != nil {
		return GetSingleSecretByIDV3Response{}, fmt.Errorf("CallGetSingleSecretByIDV3: Unable to complete api request [err=%s]", err)
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return GetSingleSecretByIDV3Response{}, ErrNotFound
		}
		return GetSingleSecretByIDV3Response{}, fmt.Errorf("CallGetSingleSecretByIDV3: Unsuccessful response. [response=%s]", response)
	}

	return secretsResponse, nil
}

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
		return GetRawSecretsV3Response{}, fmt.Errorf("CallGetSecretsRawV3: Unsuccessful response. Please make sure your secret path, workspace and environment name are all correct [response=%v]", response.RawResponse)
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

func (client Client) GetRawSecretsViaServiceToken(secretFolderPath string, envSlug string) ([]RawV3Secret, error) {
	if client.Config.ServiceToken == "" {
		return nil, fmt.Errorf("service token must be defined to fetch secrets")
	}

	serviceTokenParts := strings.SplitN(client.Config.ServiceToken, ".", 4)
	if len(serviceTokenParts) < 4 {
		return nil, fmt.Errorf("invalid service token entered. Please double check your service token and try again")
	}

	serviceTokenDetails, err := client.GetServiceTokenDetailsV2()
	if err != nil {
		return nil, fmt.Errorf("unable to get service token details. [err=%v]", err)
	}

	secrets, err := client.GetRawSecrets(secretFolderPath, envSlug, serviceTokenDetails.Workspace)

	if err != nil {
		return nil, err
	}

	return secrets, nil
}
