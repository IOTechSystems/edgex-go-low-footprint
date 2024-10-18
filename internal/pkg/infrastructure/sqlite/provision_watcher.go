//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	model "github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

// AddProvisionWatcher adds a new provision watcher
func (c *Client) AddProvisionWatcher(pw model.ProvisionWatcher) (model.ProvisionWatcher, errors.EdgeX) {
	c.loggingClient.Warn("AddProvisionWatcher function is not implemented")
	return model.ProvisionWatcher{}, nil
}

// ProvisionWatcherById gets a provision watcher by id
func (c *Client) ProvisionWatcherById(id string) (model.ProvisionWatcher, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatcherById function is not implemented")
	return model.ProvisionWatcher{}, nil
}

// ProvisionWatcherByName gets a provision watcher by name
func (c *Client) ProvisionWatcherByName(name string) (model.ProvisionWatcher, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatcherByName function is not implemented")
	return model.ProvisionWatcher{}, nil
}

// ProvisionWatchersByServiceName query provision watchers by offset, limit and service name
func (c *Client) ProvisionWatchersByServiceName(offset int, limit int, name string) ([]model.ProvisionWatcher, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatchersByServiceName function is not implemented")
	return []model.ProvisionWatcher{}, nil
}

// ProvisionWatchersByProfileName query provision watchers by offset, limit and profile name
func (c *Client) ProvisionWatchersByProfileName(offset int, limit int, name string) ([]model.ProvisionWatcher, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatchersByProfileName function is not implemented")
	return []model.ProvisionWatcher{}, nil
}

// AllProvisionWatchers query provision watchers with offset, limit and labels
func (c *Client) AllProvisionWatchers(offset int, limit int, labels []string) (pws []model.ProvisionWatcher, err errors.EdgeX) {
	c.loggingClient.Warn("AllProvisionWatchers function is not implemented")
	return []model.ProvisionWatcher{}, nil
}

// DeleteProvisionWatcherByName deletes a provision watcher by name
func (c *Client) DeleteProvisionWatcherByName(name string) errors.EdgeX {
	c.loggingClient.Warn("DeleteProvisionWatcherByName function is not implemented")
	return nil
}

// UpdateProvisionWatcher updates a provision watcher
func (c *Client) UpdateProvisionWatcher(pw model.ProvisionWatcher) errors.EdgeX {
	c.loggingClient.Warn("UpdateProvisionWatcher function is not implemented")
	return nil
}

// ProvisionWatcherCountByLabels returns the total count of Provision Watchers with labels specified.  If no label is specified, the total count of all provision watchers will be returned.
func (c *Client) ProvisionWatcherCountByLabels(labels []string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatcherCountByLabels function is not implemented")
	return 0, nil
}

// ProvisionWatcherCountByServiceName returns the count of Provision Watcher associated with specified service
func (c *Client) ProvisionWatcherCountByServiceName(name string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatcherCountByServiceName function is not implemented")
	return 0, nil
}

// ProvisionWatcherCountByProfileName returns the count of Provision Watcher associated with specified profile
func (c *Client) ProvisionWatcherCountByProfileName(name string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("ProvisionWatcherCountByProfileName function is not implemented")
	return 0, nil
}
