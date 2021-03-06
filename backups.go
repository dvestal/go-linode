package golinode

import (
	"fmt"
)

// LinodeBackup represents a linode backup
type LinodeBackup struct {
	Enabled      bool
	Availability string
	Schedule     struct {
		Day    string
		Window string
	}
	LastBackup *LinodeSnapshot
	Disks      []*LinodeInstanceDisk
}

// GetInstanceBackups gets linode backups
func (c *Client) GetInstanceBackups(linodeID int) (*LinodeInstanceBackupsResponse, error) {
	e, err := c.Instances.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d/backups", e, linodeID)
	r, err := c.R().
		SetResult(&LinodeInstanceBackupsResponse{}).
		Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LinodeInstanceBackupsResponse).fixDates(), nil
}

type LinodeBackupSnapshotResponse struct {
	Current    *LinodeSnapshot
	InProgress *LinodeSnapshot `json:"in_progress"`
}

func (l *LinodeBackupSnapshotResponse) fixDates() *LinodeBackupSnapshotResponse {
	if l.Current != nil {
		l.Current.fixDates()
	}
	if l.InProgress != nil {
		l.InProgress.fixDates()
	}
	return l
}

// LinodeInstanceBackupsResponse response struct for backup snapshot
type LinodeInstanceBackupsResponse struct {
	Automatic []*LinodeSnapshot
	Snapshot  *LinodeBackupSnapshotResponse
}

func (l *LinodeInstanceBackupsResponse) fixDates() *LinodeInstanceBackupsResponse {
	for _, el := range l.Automatic {
		el.fixDates()
	}
	if l.Snapshot != nil {
		l.Snapshot.fixDates()
	}
	return l
}
