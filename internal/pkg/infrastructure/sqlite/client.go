//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"context"

	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/sqlite"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type Client struct {
	*sqlite.Client
	loggingClient logger.LoggingClient
}

// NewClient initializes the DBClient
func NewClient(ctx context.Context, config db.Configuration, lc logger.LoggingClient) (dbClient *Client, err error) {
	dc := &Client{}
	dc.Client, err = sqlite.NewClient(ctx, config, lc)
	dc.loggingClient = lc
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "sqlite client creation failed", err)
	}
	if err := CreateTable(dc.Client.DB); err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to create SQLite table", err)
	}
	return dc, nil
}
