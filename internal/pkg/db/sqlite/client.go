//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"context"
	"database/sql"
	"sync"

	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	_ "modernc.org/sqlite"
)

var once sync.Once
var dc *Client

// Client represents an SQLite client
type Client struct {
	DB            *sql.DB
	loggingClient logger.LoggingClient
}

// NewClient returns a pointer to the SQLite client
func NewClient(ctx context.Context, config db.Configuration, lc logger.LoggingClient) (*Client, errors.EdgeX) {

	var edgeXerr errors.EdgeX
	once.Do(func() {
		// Open the SQLite database file (it will be created if it doesn't exist)
		database, err := sql.Open("sqlite", config.Host)
		if err != nil {
			edgeXerr = errors.NewCommonEdgeX(errors.KindDatabaseError, "fail to open sqlite connection", err)
			return
		}

		dc = &Client{
			DB:            database,
			loggingClient: lc,
		}
	})

	if edgeXerr != nil {
		return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to create an SQLite client", edgeXerr)
	}

	// Ping to test the connectivity to the DB
	if err := dc.DB.PingContext(ctx); err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to connect to SQLite database", err)
	}

	return dc, nil
}

// CloseSession closes the connections to SQLite
func (c *Client) CloseSession() {
	err := c.DB.Close()
	if err != nil {
		c.loggingClient.Error("error closing the SQLite database: " + err.Error())
	}
}
