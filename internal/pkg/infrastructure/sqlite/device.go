//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	stdErrs "errors"
	"fmt"

	pkgCommon "github.com/edgexfoundry/edgex-go/internal/pkg/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	model "github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/google/uuid"
)

const (
	deviceTableName = "core_metadata_device"
)

// AddDevice adds a new device
func (c *Client) AddDevice(d model.Device) (model.Device, errors.EdgeX) {
	ctx := context.Background()

	if len(d.Id) == 0 {
		d.Id = uuid.New().String()
	}

	exists, _ := deviceNameExists(ctx, c.DB, d.Name)
	if exists {
		return model.Device{}, errors.NewCommonEdgeX(errors.KindDuplicateName, fmt.Sprintf("device name %s already exists", d.Name), nil)
	}

	timestamp := pkgCommon.MakeTimestamp()
	d.Created = timestamp
	d.Modified = timestamp
	// Marshal the device to store it in the database
	deviceJSONBytes, err := json.Marshal(d)
	if err != nil {
		return model.Device{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "unable to JSON marshal device for Postgres persistence", err)
	}

	_, err = c.DB.Exec(sqlInsert(deviceTableName, idCol, contentCol), d.Id, deviceJSONBytes)
	if err != nil {
		return model.Device{}, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to insert device", err)
	}

	return d, nil
}

// DeleteDeviceById deletes a device by id
func (c *Client) DeleteDeviceById(id string) errors.EdgeX {
	c.loggingClient.Warn("DeleteDeviceById function is not implemented")
	return nil
}

// DeleteDeviceByName deletes a device by name
func (c *Client) DeleteDeviceByName(name string) errors.EdgeX {
	c.loggingClient.Warn("DeleteDeviceByName function is not implemented")
	return nil
}

// DevicesByServiceName query devices by offset, limit and name
func (c *Client) DevicesByServiceName(offset int, limit int, name string) ([]model.Device, errors.EdgeX) {
	ctx := context.Background()
	offset, validLimit := getValidOffsetAndLimit(offset, limit)
	return queryDevices(ctx, c.DB, sqlQueryContentByJSONFieldWithPagination(deviceTableName, serviceNameField), name, offset, validLimit)
}

// DeviceIdExists checks the device existence by id
func (c *Client) DeviceIdExists(id string) (bool, errors.EdgeX) {
	c.loggingClient.Warn("DeviceIdExists function is not implemented")
	return false, nil
}

// DeviceNameExists checks the device existence by name
func (c *Client) DeviceNameExists(name string) (bool, errors.EdgeX) {
	ctx := context.Background()
	return deviceNameExists(ctx, c.DB, name)
}

// DeviceById gets a device by id
func (c *Client) DeviceById(id string) (model.Device, errors.EdgeX) {
	c.loggingClient.Warn("DeviceById function is not implemented")
	return model.Device{}, nil
}

// DeviceByName gets a device by name
func (c *Client) DeviceByName(name string) (model.Device, errors.EdgeX) {
	ctx := context.Background()

	d, err := queryOneDevice(ctx, c.DB, sqlQueryContentByJSONField(deviceTableName, nameField), name)
	if err != nil {
		if stdErrs.Is(err, sql.ErrNoRows) {
			return d, errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, fmt.Sprintf("no device with name '%s' found", name), err)
		}
		return d, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to query device by name", err)
	}
	return d, nil
}

// AllDevices query the devices with offset, limit, and labels
func (c *Client) AllDevices(offset int, limit int, labels []string) (devices []model.Device, err errors.EdgeX) {
	c.loggingClient.Warn("AllDevices function is not implemented")
	return nil, nil
}

// DevicesByProfileName query devices by offset, limit and profile name
func (c *Client) DevicesByProfileName(offset int, limit int, profileName string) ([]model.Device, errors.EdgeX) {
	c.loggingClient.Warn("DevicesByProfileName function is not implemented")
	return nil, nil
}

// UpdateDevice updates a device
func (c *Client) UpdateDevice(d model.Device) errors.EdgeX {
	c.loggingClient.Warn("UpdateDevice function is not implemented")
	return nil
}

// DeviceCountByLabels returns the total count of Devices with labels specified.  If no label is specified, the total count of all devices will be returned.
func (c *Client) DeviceCountByLabels(labels []string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceCountByLabels function is not implemented")
	return 0, nil
}

// DeviceCountByProfileName returns the count of Devices associated with specified profile
func (c *Client) DeviceCountByProfileName(profileName string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceCountByProfileName function is not implemented")
	return 0, nil
}

// DeviceCountByServiceName returns the count of Devices associated with specified service
func (c *Client) DeviceCountByServiceName(serviceName string) (uint32, errors.EdgeX) {
	ctx := context.Background()
	return getTotalRowsCount(ctx, c.DB, sqlQueryCountByJSONField(deviceTableName, serviceNameField), serviceName)
}

func deviceNameExists(ctx context.Context, db *sql.DB, name string) (bool, errors.EdgeX) {
	var exists bool
	err := db.QueryRowContext(ctx, sqlCheckExistsByJSONField(deviceTableName, nameField), name).Scan(&exists)
	if err != nil {
		return false, errors.NewCommonEdgeX(errors.KindDatabaseError, fmt.Sprintf("failed to query device by name '%s' from %s table", name, deviceTableName), err)
	}
	return exists, nil
}

func queryDevices(ctx context.Context, db *sql.DB, sql string, args ...any) ([]model.Device, errors.EdgeX) {
	rows, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to query devices", err)
	}
	defer rows.Close()

	var devices []model.Device
	for rows.Next() {
		var content []byte
		if err = rows.Scan(&content); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to scan device content", err)
		}

		var d model.Device
		if err = json.Unmarshal(content, &d); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to unmarshal device", err)
		}
		devices = append(devices, d)
	}

	return devices, nil
}

func queryOneDevice(ctx context.Context, db *sql.DB, sql string, args ...any) (model.Device, error) {
	var d model.Device
	row := db.QueryRowContext(ctx, sql, args...)

	var content []byte
	if err := row.Scan(&content); err != nil {
		return d, err
	}

	if err := json.Unmarshal(content, &d); err != nil {
		return d, fmt.Errorf("failed to unmarshal device JSON: %w", err)
	}

	return d, nil
}
