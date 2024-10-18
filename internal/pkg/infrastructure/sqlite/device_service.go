//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	pkgCommon "github.com/edgexfoundry/edgex-go/internal/pkg/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	model "github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/google/uuid"
)

const (
	deviceServiceTableName = "core_metadata_device_service"
)

// AddDeviceService adds a new device service
func (c *Client) AddDeviceService(ds model.DeviceService) (model.DeviceService, errors.EdgeX) {
	ctx := context.Background()

	if len(ds.Id) == 0 {
		ds.Id = uuid.New().String()
	}

	// verify if device service name is unique or not
	exists, edgexErr := deviceServiceNameExists(ctx, c.DB, ds.Name)
	if edgexErr != nil {
		return model.DeviceService{}, errors.NewCommonEdgeX(errors.KindDatabaseError, "device service name existence check failed", edgexErr)
	}
	if exists {
		return model.DeviceService{}, errors.NewCommonEdgeX(errors.KindDuplicateName, fmt.Sprintf("device service name %s already exists", ds.Name), nil)
	}

	timestamp := pkgCommon.MakeTimestamp()
	ds.Created = timestamp
	ds.Modified = timestamp

	deviceServiceJSONBytes, err := json.Marshal(ds)
	if err != nil {
		return model.DeviceService{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "unable to JSON marshal device service for SQLite persistence", err)
	}

	_, err = c.DB.ExecContext(ctx, sqlInsert(deviceServiceTableName, idCol, contentCol), ds.Id, string(deviceServiceJSONBytes))
	if err != nil {
		return model.DeviceService{}, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to insert device service: ", err)
	}

	return ds, nil
}

// DeviceServiceById gets a device service by id
func (c *Client) DeviceServiceById(id string) (deviceService model.DeviceService, edgeXerr errors.EdgeX) {
	c.loggingClient.Warn("DeviceServiceById function is not implemented")
	return model.DeviceService{}, nil
}

// DeviceServiceByName gets a device service by name
func (c *Client) DeviceServiceByName(name string) (deviceService model.DeviceService, edgeXerr errors.EdgeX) {
	ctx := context.Background()
	ds, err := queryOneDeviceService(ctx, c.DB, sqlQueryContentByJSONField(deviceServiceTableName, nameField), name)
	if err != nil {
		if err == sql.ErrNoRows {
			return ds, errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, fmt.Sprintf("no device service with name '%s' found", name), err)
		}
		return ds, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to query device service by name", err)
	}
	return ds, nil
}

// DeleteDeviceServiceById deletes a device service by id
func (c *Client) DeleteDeviceServiceById(id string) errors.EdgeX {
	c.loggingClient.Warn("DeleteDeviceServiceById function is not implemented")
	return nil
}

// DeleteDeviceServiceByName deletes a device service by name
func (c *Client) DeleteDeviceServiceByName(name string) errors.EdgeX {
	c.loggingClient.Warn("DeleteDeviceServiceByName function is not implemented")
	return nil
}

// DeviceServiceNameExists checks the device service exists by name
func (c *Client) DeviceServiceNameExists(name string) (bool, errors.EdgeX) {
	ctx := context.Background()
	return deviceServiceNameExists(ctx, c.DB, name)
}

// AllDeviceServices returns multiple device services per query criteria, including
// offset: the number of items to skip before starting to collect the result set
// limit: The numbers of items to return
// labels: allows for querying a given object by associated user-defined labels
func (c *Client) AllDeviceServices(offset int, limit int, labels []string) (deviceServices []model.DeviceService, err errors.EdgeX) {
	c.loggingClient.Warn("AllDeviceServices function is not implemented")
	return nil, nil
}

// UpdateDeviceService updates a device service
func (c *Client) UpdateDeviceService(ds model.DeviceService) errors.EdgeX {
	c.loggingClient.Warn("UpdateDeviceService function is not implemented")
	return nil
}

// DeviceServiceCountByLabels returns the total count of Device Services with labels specified.  If no label is specified, the total count of all device services will be returned.
func (c *Client) DeviceServiceCountByLabels(labels []string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceServiceCountByLabels function is not implemented")
	return 0, nil
}

func deviceServiceNameExists(ctx context.Context, db *sql.DB, name string) (bool, errors.EdgeX) {
	var exists bool
	err := db.QueryRowContext(ctx, sqlCheckExistsByJSONField(deviceServiceTableName, nameField), name).Scan(&exists)
	if err != nil {
		return false, errors.NewCommonEdgeX(errors.KindDatabaseError, fmt.Sprintf("failed to query device service by name '%s' from %s table", name, deviceServiceTableName), err)
	}
	return exists, nil
}

func queryOneDeviceService(ctx context.Context, db *sql.DB, sql string, args ...any) (model.DeviceService, error) {
	var ds model.DeviceService
	row := db.QueryRowContext(ctx, sql, args...)

	var content []byte
	if err := row.Scan(&content); err != nil {
		return ds, err
	}

	if err := json.Unmarshal(content, &ds); err != nil {
		return ds, fmt.Errorf("failed to unmarshal device profile JSON: %w", err)
	}

	return ds, nil
}
