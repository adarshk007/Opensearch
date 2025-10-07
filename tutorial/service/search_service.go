package service

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/adarshk007/tutorial/client"
)

type SearchService struct{}

// Search all documents in an index
func (s *SearchService) SearchAll(index string) (map[string]interface{}, error) {
	query := `{"query":{"match_all":{}}}`

	res, err := client.OSClient.Client.Search(
		client.OSClient.Client.Search.WithContext(context.Background()),
		client.OSClient.Client.Search.WithIndex(index),
		client.OSClient.Client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// You can extend: Index, BulkIndex, Delete, Update
