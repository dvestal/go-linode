package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	golinode "github.com/chiefy/go-linode"
	"github.com/dnaeon/go-vcr/recorder"
)

func main() {

	envErr := fmt.Errorf("env vars LINODE_INSTANCE_ID and LINODE_VOLUME_ID must be set")

	var linodeInstanceID int
	var linodeVolumeID int
	var err error

	if linodeInstanceID, err = strconv.Atoi(os.Getenv("LINODE_INSTANCE_ID")); err != nil {
		log.Fatal(envErr)
	}
	if linodeVolumeID, err = strconv.Atoi(os.Getenv("LINODE_VOLUME_ID")); err != nil {
		log.Fatal(envErr)
	}

	// Start our recorder
	r, err := recorder.New("test/fixtures")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	c, err := golinode.NewClient(nil, r)
	if err != nil {
		log.Fatalf("Failed to create linode client: %s", err)
	}
	c.SetDebug(false)

	_, err = c.ListRegions()
	if err != nil {
		log.Fatalf("Failed to get linode regions: %s", err)
	}
	log.Println("Succesfully got linode regions")

	_, err = c.ListInstances(nil)
	if err != nil {
		log.Fatalf("Failed to get linode instances: %s", err)
	}
	log.Println("Succesfully got linode instances")

	_, err = c.ListImages()
	if err != nil {
		log.Fatalf("Failed to get linode images: %s", err)
	}
	log.Println("Succesfully got linode images")

	_, err = c.GetInstance(linodeInstanceID)
	if err != nil {
		log.Fatalf("Failed to get linode instance ID %d: %s", linodeInstanceID, err)
	}
	log.Println(fmt.Sprintf("Succesfully got linode instance ID %d", linodeInstanceID))

	_, err = c.GetInstanceBackups(linodeInstanceID)
	if err != nil {
		log.Fatalf("Failed to get linode backups for instance ID %d: %s", linodeInstanceID, err)
	}
	log.Println(fmt.Sprintf("Succesfully got linode backups for instance ID %d", linodeInstanceID))

	_, err = c.ListInstanceDisks(linodeInstanceID)
	if err != nil {
		log.Fatalf("Failed to get linode instance disks: %s", err)
	}
	log.Println("Succesfully got linode instance disks")

	_, err = c.ListInstanceConfigs(linodeInstanceID)
	if err != nil {
		log.Fatalf("Failed to get linode instance configs: %s", err)
	}
	log.Println("Succesfully got linode instance configs")

	_, err = c.ListInstanceVolumes(linodeInstanceID)
	if err != nil {
		log.Fatalf("Failed to get linode instance volumes: %s", err)
	}
	log.Println("Succesfully got linode instance volumes")

	_, err = c.ListStackscripts()
	if err != nil {
		log.Fatalf("Failed to get linode stackscripts: %s", err)
	}
	log.Println("Succesfully got linode public stackscripts (1 page)")

	_, err = c.GetStackscript(7)
	if err != nil {
		log.Fatalf("Failed to get linode stackscript ID 7: %s", err)
	}
	log.Println("Succesfully got linode stackscript ID 7")

	_, err = c.ListVolumes()
	if err != nil {
		log.Fatalf("Failed to get linode volumes: %s", err)
	}
	log.Println("Succesfully got linode volumes (1 page)")

	_, err = c.GetVolume(linodeVolumeID)
	if err != nil {
		log.Fatalf("Failed to get linode volume ID %d: %s", linodeVolumeID, err)
	}
	log.Println(fmt.Sprintf("Succesfully got linode volume ID %d", linodeVolumeID))

	log.Printf("Successfully retrieved linode requests!")
}
