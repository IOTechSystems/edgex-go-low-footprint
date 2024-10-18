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
	deviceProfileTableName = "core_metadata_device_profile"
)

func (c *Client) AddDeviceProfile(dp model.DeviceProfile) (model.DeviceProfile, errors.EdgeX) {
	ctx := context.Background()

	if len(dp.Id) == 0 {
		dp.Id = uuid.New().String()
	}

	// verify if device profile name is unique or not
	exists, err := deviceProfileNameExists(ctx, c.DB, dp.Name)
	if err != nil {
		return model.DeviceProfile{}, errors.NewCommonEdgeX(errors.KindDatabaseError, "device profile name existence check failed", err)
	}
	if exists {
		return model.DeviceProfile{}, errors.NewCommonEdgeX(errors.KindDuplicateName, fmt.Sprintf("device profile name %s already exists", dp.Name), nil)
	}

	timestamp := pkgCommon.MakeTimestamp()
	dp.Created = timestamp
	dp.Modified = timestamp

	deviceProfileJSONBytes, err := json.Marshal(dp)
	if err != nil {
		return model.DeviceProfile{}, errors.NewCommonEdgeX(errors.KindContractInvalid, "unable to JSON marshal device profile for Postgres persistence", err)
	}

	_, err = c.DB.ExecContext(ctx, sqlInsert(deviceProfileTableName, idCol, contentCol), dp.Id, string(deviceProfileJSONBytes))
	if err != nil {
		return model.DeviceProfile{}, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to insert device profile: ", err)
	}

	return dp, nil
}

func (c *Client) UpdateDeviceProfile(dp model.DeviceProfile) errors.EdgeX {
	c.loggingClient.Warn("UpdateDeviceProfile function is not implemented")
	return nil
}

func (c *Client) DeviceProfileById(id string) (model.DeviceProfile, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfileById function is not implemented")
	return model.DeviceProfile{}, nil
}

func (c *Client) DeviceProfileByName(name string) (model.DeviceProfile, errors.EdgeX) {
	ctx := context.Background()
	dp, err := queryOneDeviceProfile(ctx, c.DB, sqlQueryContentByJSONField(deviceProfileTableName, nameField), name)
	if err != nil {
		if err == sql.ErrNoRows {
			return dp, errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, fmt.Sprintf("no device profile with name '%s' found", name), err)
		}
		return dp, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to query device profile by name", err)
	}
	return dp, nil
}

func (c *Client) DeleteDeviceProfileById(id string) errors.EdgeX {
	c.loggingClient.Warn("DeleteDeviceProfileById function is not implemented")
	return nil
}

func (c *Client) DeleteDeviceProfileByName(name string) errors.EdgeX {
	c.loggingClient.Warn("DeleteDeviceProfileByName function is not implemented")
	return nil
}

func (c *Client) DeviceProfileNameExists(name string) (bool, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfileNameExists function is not implemented")
	return false, nil
}

func (c *Client) AllDeviceProfiles(offset int, limit int, labels []string) (profiles []model.DeviceProfile, err errors.EdgeX) {
	c.loggingClient.Warn("AllDeviceProfiles function is not implemented")
	return nil, nil
}

func (c *Client) DeviceProfilesByModel(offset int, limit int, model string) ([]model.DeviceProfile, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfilesByModel function is not implemented")
	return nil, nil
}

func (c *Client) DeviceProfilesByManufacturer(offset int, limit int, manufacturer string) ([]model.DeviceProfile, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfilesByManufacturer function is not implemented")
	return nil, nil
}

func (c *Client) DeviceProfilesByManufacturerAndModel(offset int, limit int, manufacturer string, model string) ([]model.DeviceProfile, uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfilesByManufacturerAndModel function is not implemented")
	return nil, 0, nil
}

func (c *Client) DeviceProfileCountByLabels(labels []string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfileCountByLabels function is not implemented")
	return 0, nil
}

func (c *Client) DeviceProfileCountByManufacturer(manufacturer string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfileCountByManufacturer function is not implemented")
	return 0, nil
}

func (c *Client) DeviceProfileCountByModel(model string) (uint32, errors.EdgeX) {
	c.loggingClient.Warn("DeviceProfileCountByModel function is not implemented")
	return 0, nil
}

func deviceProfileNameExists(ctx context.Context, db *sql.DB, name string) (bool, error) {
	var exists bool
	err := db.QueryRowContext(ctx, sqlCheckExistsByJSONField(deviceProfileTableName, nameField), name).Scan(&exists)
	if err != nil {
		return false, errors.NewCommonEdgeX(errors.KindDatabaseError, fmt.Sprintf("failed to query device profile by name '%s' from %s table", name, deviceProfileTableName), err)
	}
	return exists, nil
}

func queryOneDeviceProfile(ctx context.Context, db *sql.DB, sql string, args ...any) (model.DeviceProfile, error) {
	var dp model.DeviceProfile
	row := db.QueryRowContext(ctx, sql, args...)

	var content []byte
	if err := row.Scan(&content); err != nil {
		return dp, err
	}

	if err := json.Unmarshal(content, &dp); err != nil {
		return dp, fmt.Errorf("failed to unmarshal device profile JSON: %w", err)
	}

	return dp, nil
}
