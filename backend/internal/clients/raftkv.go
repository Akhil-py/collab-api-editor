package clients

import (
    "github.com/go-resty/resty/v2"
    "os"
)

type RaftKVClient struct {
    client *resty.Client
}

func NewRaftKVClient() *RaftKVClient {
    return &RaftKVClient{
        client: resty.New().SetBaseURL(os.Getenv("RAFT_KV_URL")),
    }
}

// Example: Get session metadata
func (r *RaftKVClient) GetSession(key string) (string, error) {
    resp, err := r.client.R().Get("/session/" + key)
    if err != nil {
        return "", err
    }
    return resp.String(), nil
} 