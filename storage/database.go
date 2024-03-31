package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"time"
)

const (
	tableName string = `accounts`

	createSQLPrefix string = `CREATE TABLE IF NOT EXISTS`
	createSQLSuffix string = `(
	  id INTEGER NOT NULL PRIMARY KEY,
	  time DATETIME NOT NULL,
	  label TEXT NOT NULL UNIQUE,
	  issuer TEXT NOT NULL,
	  secret TEXT NOT NULL,
      digits INTEGER,
      period INTEGER,
	  algorithm TEXT NOT NULL
  );`

	insertSQLPrefix string = `INSERT INTO`
	insertSQLSuffix string = `(time, label, issuer, secret, digits, period, algorithm) VALUES(?, ?, ?, ?, ?, ?, ?);`

	updateSQLPrefix string = `UPDATE`
	updateSQLSuffix string = `SET label = ?, issuer = ?, secret = ?, digits = ?, period = ?, algorithm = ? WHERE id = ?;`

	selectSQLPrefix string = `SELECT id, time, label, issuer, secret, digits, period, algorithm FROM`
)

type Database struct {
	fileName      string
	database      *sql.DB
	insertCommand *sql.Stmt
	updateCommand *sql.Stmt
}

func sqlCommand(prefix string, table string, suffix string) string {
	return strings.TrimSpace(prefix + " " + table + " " + suffix)
}

func Open(fileName string) (*Database, error) {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(sqlCommand(createSQLPrefix, tableName, createSQLSuffix)); err != nil {
		return nil, err
	}

	insertCommand, err := db.Prepare(sqlCommand(insertSQLPrefix, tableName, insertSQLSuffix))
	if err != nil {
		return nil, err
	}

	updateCommand, err := db.Prepare(sqlCommand(updateSQLPrefix, tableName, updateSQLSuffix))
	if err != nil {
		return nil, err
	}

	return &Database{fileName: fileName, database: db, insertCommand: insertCommand, updateCommand: updateCommand}, nil
}

func (database *Database) Close() error {
	return database.database.Close()
}

func (database *Database) Insert(label string, issuer string, secret string, digits uint8, period uint8, algorithm string) error {
	_, err := database.insertCommand.Exec(time.Now().UTC(), label, issuer, secret, digits, period, algorithm)
	return err
}

func (database *Database) Update(record Record) error {
	_, err := database.updateCommand.Exec(record.Label, record.Issuer, record.Secret, record.Digits, record.Period, record.Algorithm, record.Id)
	return err
}

func (database *Database) Query() (Records, error) {
	rows, err := database.database.Query(sqlCommand(selectSQLPrefix, tableName, ""))
	if err != nil {
		return Records{}, err
	}
	defer rows.Close()

	var dataList Records
	for rows.Next() {
		var data Record
		if err := rows.Scan(&data.Id, &data.Time, &data.Label, &data.Issuer, &data.Secret, &data.Digits, &data.Period, &data.Algorithm); err != nil {
			return Records{}, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}
