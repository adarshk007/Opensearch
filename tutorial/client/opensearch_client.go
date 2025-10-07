package client

import (
	"crypto/tls"
	"log"
	"net/http"
	"sync"

	"github.com/adarshk007/tutorial/config"
	"github.com/opensearch-project/opensearch-go/v2"
)

type OpenSearchClient struct {
	Client *opensearch.Client
}

var (
	OSClient *OpenSearchClient
	once     sync.Once
)

// InitOpenSearch initializes singleton client
func InitOpenSearch() {
	once.Do(func() {
		cfg := opensearch.Config{
			Addresses: []string{config.Cfg.OpenSearch.URL},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Username: config.Cfg.OpenSearch.Username,
			Password: config.Cfg.OpenSearch.Password,
		}

		client, err := opensearch.NewClient(cfg)
		if err != nil {
			log.Fatalf("Error creating OpenSearch client: %s", err)
		}

		OSClient = &OpenSearchClient{Client: client}
	})
}
