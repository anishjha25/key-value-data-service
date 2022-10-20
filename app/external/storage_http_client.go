package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"key-value-data-service/app/dto"
	"log"
	"net/http"
)

type StorageHTTPClient interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type storageHTTPClient struct {
	host   string
	client *http.Client
}

func (s *storageHTTPClient) Set(key, value string) error {
	body := dto.KeyValue{
		Value: value,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("error marshal request: %v", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", s.host, key), bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Printf("error creating http request: %v", err)
		return err
	}

	_, err = s.client.Do(req)
	if err != nil {
		log.Printf("error: %v in setting key: %s", err, key)
	}

	return err
}

func (s *storageHTTPClient) Get(key string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", s.host, key), nil)
	if err != nil {
		log.Printf("error creating http request: %v", err)
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("error: %v in getting key: %s", err, key)
		return "", err
	}

	if resp.StatusCode == http.StatusNotFound {
		return "", fmt.Errorf("not found")
	}

	defer resp.Body.Close()

	var result dto.KeyValue
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("error: %v in decoding  response for key: %s", err, key)
		return "", err
	}

	return result.Value, nil

}

func (s *storageHTTPClient) Delete(key string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", s.host, key), nil)
	if err != nil {
		log.Printf("error creating http request: %v", err)
		return err
	}

	_, err = s.client.Do(req)
	if err != nil {
		log.Printf("error: %v in deleting key: %s", err, key)
	}

	return err
}

func NewStorageHTTPClient(host string) StorageHTTPClient {
	return &storageHTTPClient{
		client: http.DefaultClient,
		host:   host,
	}
}
