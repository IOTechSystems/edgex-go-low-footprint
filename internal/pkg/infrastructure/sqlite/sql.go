//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package sqlite

import (
	"database/sql"
	"fmt"
	"strings"
)

// sqlInsert returns the SQL statement for inserting a new row with the given columns into the table.
func sqlInsert(table string, columns ...string) string {
	columnCount := len(columns)
	columnNames := strings.Join(columns, ", ")
	valuePlaceholders := make([]string, columnCount)

	// SQLite uses '?' placeholders
	for i := 0; i < columnCount; i++ {
		valuePlaceholders[i] = "?"
	}

	valueNames := strings.Join(valuePlaceholders, ", ")
	return fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", table, columnNames, valueNames)
}

// sqlQueryContentByJSONField returns the SQL statement for selecting content column in the table by the given JSON query string
func sqlQueryContentByJSONField(table string, fieldName string) string {
	// Query for JSON field 'name' in the 'content' column
	return fmt.Sprintf("SELECT content FROM %s WHERE json_extract(content, '$.%s') = ?", table, fieldName)
}

// sqlQueryContentByJSONFieldWithPagination returns the SQL statement for selecting content column in the table by the given JSON query string with pagination
func sqlQueryContentByJSONFieldWithPagination(table string, fieldName string) string {
	return fmt.Sprintf("SELECT content FROM %s WHERE json_extract(content, '$.%s') = ? ORDER BY COALESCE(CAST(json_extract(content, '$.%s') AS INTEGER), 0) LIMIT ? OFFSET ?", table, fieldName, createdField)
}

// // sqlCheckExistsByJSONField returns the SQL statement for checking if a row exists by query the JSON field in content column.
func sqlCheckExistsByJSONField(table string, fieldName string) string {
	return fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE json_extract(content, '$.%s') = ?)", table, fieldName)
}

// sqlQueryCountByJSONField returns the SQL statement for counting the number of rows in the table by the given JSON query string
func sqlQueryCountByJSONField(table string, fieldName string) string {
	return fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE json_extract(content, '$.%s') = ?", table, fieldName)
}

// CreateTable creates the core metadata tables if they do not exist
func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS core_metadata_device_service (
			id TEXT PRIMARY KEY,
			content TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS core_metadata_device_profile (
			id TEXT PRIMARY KEY,
			content TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS core_metadata_device (
			id TEXT PRIMARY KEY,
			content TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS core_metadata_provision_watcher (
			id TEXT PRIMARY KEY,
			content TEXT NOT NULL
		);
	`)
	return err
}
