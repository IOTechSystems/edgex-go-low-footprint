//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"context"
	"database/sql"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

// getValidOffsetAndLimit returns the valid or default offset and limit from the given parameters
// Note: the returned limit can be an integer or nil (which means that clients want to retrieve all remaining rows after offset)
func getValidOffsetAndLimit(offset, limit int) (int, any) {
	defaultOffset := 0
	if offset < 0 {
		offset = defaultOffset
	}

	// Since PostgreSQL does not support negative limit, we need to set the default limit to nil,
	// nil limit means that clients want to retrieve all remaining rows after offset from the DB
	var defaultLimit any = nil
	if limit < 0 {
		return offset, defaultLimit
	} else {
		return offset, limit
	}
}

// getTotalRowsCount returns the total rows count from the given sql query
// Note: the sql query must be a count query, e.g. SELECT COUNT(*) FROM table_name
func getTotalRowsCount(ctx context.Context, db *sql.DB, sql string, args ...any) (uint32, errors.EdgeX) {
	var rowCount int
	err := db.QueryRowContext(ctx, sql, args...).Scan(&rowCount)
	if err != nil {
		return 0, errors.NewCommonEdgeX(errors.KindDatabaseError, "failed to query total rows count", err)
	}

	return uint32(rowCount), nil
}
