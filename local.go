package cloudconfigclient

import (
	"fmt"
	"github.com/Piszmog/cfservices"
	"net/http"
	"os"
	"strings"
)

const (
	// EnvironmentLocalConfigServerUrls is an environment variable for setting base URLs for local Config Servers.
	EnvironmentLocalConfigServerUrls = "CONFIG_SERVER_URLS"
)

// NewLocalClientFromEnv creates a ConfigClient for a locally running Config Server. Acquires the base URLs from the
// environment variable 'CONFIG_SERVER_URLS'.
//
// The ConfigClient's underlying http.Client is configured with timeouts and connection pools.
func NewLocalClientFromEnv(client *http.Client) (*ConfigClient, error) {
	serviceCredentials, err := GetLocalCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to create a local client: %w", err)
	}
	baseUrls := make([]string, len(serviceCredentials.Credentials), len(serviceCredentials.Credentials))
	for index, cred := range serviceCredentials.Credentials {
		baseUrls[index] = cred.Uri
	}
	return NewLocalClient(client, baseUrls)
}

// NewLocalClient creates a ConfigClient for a locally running Config Server.
//
// The ConfigClient's underlying http.Client is configured with timeouts and connection pools.
func NewLocalClient(client *http.Client, baseUrls []string) (*ConfigClient, error) {
	configClients := make([]CloudClient, len(baseUrls), len(baseUrls))
	for index, baseUrl := range baseUrls {
		configUri := baseUrl
		configClients[index] = Client{ConfigUri: configUri, HttpClient: client}
	}
	return &ConfigClient{Clients: configClients}, nil
}

// GetLocalCredentials creates the credentials that are used to configure a ConfigClient to access a local Config Server.
//
// Retrieves the base URLs of Config Servers from the environment variable 'CONFIG_SERVER_URLS' - a comma separated list.
func GetLocalCredentials() (*cfservices.ServiceCredentials, error) {
	localUrls := os.Getenv(EnvironmentLocalConfigServerUrls)
	if len(localUrls) == 0 {
		return nil, fmt.Errorf("no local Config Server URLs provided in environment variable %s", EnvironmentLocalConfigServerUrls)
	}
	urls := strings.Split(localUrls, ",")
	creds := make([]cfservices.Credentials, len(urls), len(urls))
	for index, url := range urls {
		creds[index] = cfservices.Credentials{Uri: url}
	}
	return &cfservices.ServiceCredentials{Credentials: creds}, nil
}
