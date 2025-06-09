package infisicalclient

import (
	"time"
)

type GetServiceTokenDetailsResponse struct {
	ID           string    `json:"_id"`
	Name         string    `json:"name"`
	Workspace    string    `json:"workspace"`
	Environment  string    `json:"environment"`
	ExpiresAt    time.Time `json:"expiresAt"`
	EncryptedKey string    `json:"encryptedKey"`
	Iv           string    `json:"iv"`
	Tag          string    `json:"tag"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	V            int       `json:"__v"`
}

type MachineIdentityAuthResponse struct {
	AccessToken       string `json:"accessToken"`
	ExpiresIn         int    `json:"expiresIn"`
	AccessTokenMaxTTL int    `json:"accessTokenMaxTTL"`
	TokenType         string `json:"tokenType"`
}

type GetRawSecretsV3Request struct {
	Environment            string `json:"environment"`
	WorkspaceId            string `json:"workspaceId"`
	SecretPath             string `json:"secretPath"`
	ExpandSecretReferences bool   `json:"expandSecretReferences"`
}

type RawV3Secret struct {
	Version       int    `json:"version"`
	Workspace     string `json:"workspace"`
	Type          string `json:"type"`
	Environment   string `json:"environment"`
	SecretKey     string `json:"secretKey"`
	SecretValue   string `json:"secretValue"`
	SecretComment string `json:"secretComment"`

	SecretReminderNote       string `json:"secretReminderNote"`
	SecretReminderRepeatDays int64  `json:"secretReminderRepeatDays"`
	Tags                     []struct {
		ID    string `json:"id"`
		Slug  string `json:"slug"`
		Color string `json:"color"`
		Name  string `json:"name"`
	} `json:"tags"`
}

type GetRawSecretsV3Response struct {
	Secrets []RawV3Secret `json:"secrets"`
}
