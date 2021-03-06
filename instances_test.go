package golinode

import (
	"testing"
)

const TestInstanceID = 7256733

func TestListInstances(t *testing.T) {
	client, err := createTestClient(debugAPI)
	if err != nil {
		t.Errorf("Error creating test client %v", err)
	}
	linodes, err := client.ListInstances(nil)
	if err != nil {
		t.Errorf("Error listing instances, expected struct, got error %v", err)
	}
	if len(linodes) != 1 {
		t.Errorf("Expected a list of instances, but got %v", linodes)
	}
}

func TestGetInstance(t *testing.T) {
	client, err := createTestClient(debugAPI)
	if err != nil {
		t.Errorf("Error creating test client %v", err)
	}
	instance, err := client.GetInstance(TestInstanceID)
	if err != nil {
		t.Errorf("Error getting instance TestInstanceID, expected *LinodeInstance, got error %v", err)
	}
	if instance.Specs.Disk <= 0 {
		t.Errorf("Error in instance TestInstanceID spec for disk size, %v", instance.Specs)
	}
}

func TestListInstanceDisks(t *testing.T) {
	client, err := createTestClient(debugAPI)
	if err != nil {
		t.Errorf("Error creating test client %v", err)
	}
	disks, err := client.ListInstanceDisks(TestInstanceID)
	if err != nil {
		t.Errorf("Error listing instance disks, expected struct, got error %v", err)
	}
	if len(disks) == 0 {
		t.Errorf("Expected a list of instance disks, but got %v", disks)
	}
}

func TestListInstanceConfigs(t *testing.T) {
	client, err := createTestClient(debugAPI)
	if err != nil {
		t.Errorf("Error creating test client %v", err)
	}
	configs, err := client.ListInstanceConfigs(TestInstanceID)
	if err != nil {
		t.Errorf("Error listing instance configs, expected struct, got error %v", err)
	}
	if len(configs) == 0 {
		t.Errorf("Expected a list of instance configs, but got %v", configs)
	}
}

func TestListInstanceVolumes(t *testing.T) {
	client, err := createTestClient(debugAPI)
	if err != nil {
		t.Errorf("Error creating test client %v", err)
	}
	volumes, err := client.ListInstanceVolumes(TestInstanceID)
	if err != nil {
		t.Errorf("Error listing instance volumes, expected struct, got error %v", err)
	}
	if len(volumes) == 0 {
		t.Errorf("Expected a list of instance volumes, but got %v", volumes)
	}
}
