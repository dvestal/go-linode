package golinode

import (
	"fmt"
	"testing"
)

func TestResourceEndpoint(t *testing.T) {
	apiKey := "MYFAKEAPIKEY"
	client, err := NewClient(&apiKey, nil)
	if err != nil {
		t.Error("Could not create new client in test")
	}
	r := client.Resource("images")
	e, err := r.Endpoint()
	if err != nil {
		t.Error("Got error when querying for images endpoint")
	}
	if e != imagesEndpoint {
		t.Errorf("Images endpoint did not match '%s'", imagesEndpoint)
	}
}
func TestResourceTemplatedEndpointWithID(t *testing.T) {
	apiKey := "MYFAKEAPIKEY"
	client, err := NewClient(&apiKey, nil)
	backupID := 1234255
	e, err := client.Backups.EndpointWithID(backupID)
	if err != nil {
		t.Error("Got error when getting endpoint with id for backups")
	}
	if e != fmt.Sprintf("linode/instances/%d/backups", backupID) {
		t.Errorf("Backups endpoint did not contain backup ID '%d'", backupID)
	}
}
