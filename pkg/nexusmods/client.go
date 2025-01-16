package nexusmods

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"time"
)

type NexusModsClient struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string
}

// TODO: Add proper rate limiting https://app.swaggerhub.com/apis-docs/NexusMods/nexus-mods_public_api_params_in_form_data/1.0#/

func NewNexusModsClient(apiKey string) *NexusModsClient {
	return &NexusModsClient{
		BaseURL: "https://api.nexusmods.com/v1",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		APIKey: apiKey,
	}
}

func (client *NexusModsClient) GenerateDownloadLink(gameDomainName string, modID, fileID string, key string, expires string) (*DownloadLinkResponse, error) {
	url := fmt.Sprintf("%s/games/%s/mods/%s/files/%s/download_link.json", client.BaseURL, gameDomainName, modID, fileID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if len(key) != 0 && len(expires) != 0 {
		q.Add("key", key)
		q.Add("expires", expires)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("apikey", client.APIKey)

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	parsedResponse, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, ErrInvalidKeyOrExpireTime
	case http.StatusForbidden:
		return nil, ErrPermissionDenied
	case http.StatusNotFound:
		return nil, ErrFileNotFound
	case http.StatusGone:
		return nil, ErrLinkExpired
	case http.StatusOK:
		return parsedResponse, nil
	default:
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}
}

func parseResponse(resp *http.Response) (*DownloadLinkResponse, error) {
	var downloadLinkResponses []DownloadLinkResponse

	err := json.NewDecoder(resp.Body).Decode(&downloadLinkResponses)
	if err != nil {
		return nil, err
	}

	if len(downloadLinkResponses) > 0 {
		return &downloadLinkResponses[0], nil
	}
	return nil, errors.New("empty response body")
}
